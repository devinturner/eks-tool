package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/eks/eksiface"
)

var REGIONS []string = []string{
	"eu-north-1",
	"ap-south-1",
	"eu-west-3",
	"eu-west-2",
	"eu-west-1",
	"ap-northeast-3",
	"ap-northeast-2",
	"ap-northeast-1",
	"sa-east-1",
	"ca-central-1",
	"ap-southeast-1",
	"ap-southeast-2",
	"eu-central-1",
	"us-east-1",
	"us-east-2",
	"us-west-1",
	"us-west-2",
}

type Client struct {
	eksiface.EKSAPI
}

func main() {
	rs := make([]*Region, 0)
	for _, region := range REGIONS {
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Config:            aws.Config{Region: aws.String(region)},
		}))

		client := &Client{eks.New(sess)}

		clusters, err := client.GetClusters()
		if err != nil {
			log.Fatal(err)
		}

		if len(clusters) >= 1 {
			rs = append(rs, &Region{
				Name:     region,
				Clusters: clusters,
			})
		}
	}

	out, err := json.Marshal(rs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

}

type Region struct {
	Name     string      `json:"name"`
	Clusters ClusterList `json:"clusters"`
}

type ClusterList []*Cluster

type Cluster struct {
	Name       string   `json:"name"`
	NodeGroups []string `json:"node_groups"`
}

func (c *Client) GetNodeGroups(clusterName string) ([]string, error) {
	out := make([]string, 0)
	res, err := c.ListNodegroups(&eks.ListNodegroupsInput{
		ClusterName: aws.String(clusterName),
	})
	if err != nil {
		return nil, err
	}

	for _, ng := range res.Nodegroups {
		out = append(out, *ng)
	}
	return out, nil
}

func (c *Client) GetClusters() (ClusterList, error) {
	clusters := make(ClusterList, 0)
	res, err := c.ListClusters(&eks.ListClustersInput{})
	if err != nil {
		return nil, err
	}

	for _, cc := range res.Clusters {
		ngs, err := c.GetNodeGroups(*cc)
		if err != nil {
			return nil, err
		}
		if len(ngs) >= 1 {
			clusters = append(clusters, &Cluster{
				Name:       *cc,
				NodeGroups: ngs,
			})
		}
	}

	return clusters, nil
}
