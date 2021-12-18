package cloud

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
)

type NodeGroup struct {
	Name      string `json:"name"`
	NodeCount int    `json:"node_count"`
}

func (c *Client) GetNodeGroups(cluster string) ([]*NodeGroup, error) {
	ngs := make([]*NodeGroup, 0)
	res, err := c.ListNodegroups(&eks.ListNodegroupsInput{
		ClusterName: aws.String(cluster),
	})
	if err != nil {
		return nil, err
	}

	for _, ng := range res.Nodegroups {
		ngs = append(ngs, &NodeGroup{Name: *ng})
	}
	//for _, ng := range res.Nodegroups {
	//	n, err := c.GetNodeGroup(cluster, *ng)
	//	if err != nil {
	//		return nil, err
	//	}
	//	ngs = append(ngs, &NodeGroup{Name: n.Name, NodeCount: n.NodeCount})
	//}
	return ngs, nil
}

func (c *Client) GetNodeGroup(cluster, name string) (*NodeGroup, error) {
	res, err := c.DescribeNodegroup(&eks.DescribeNodegroupInput{
		ClusterName:   aws.String(cluster),
		NodegroupName: aws.String(name),
	})
	if err != nil {
		return nil, err
	}
	ng := new(NodeGroup)
	ng.Name = *res.Nodegroup.NodegroupName
	ng.NodeCount = int(*res.Nodegroup.ScalingConfig.DesiredSize)
	return ng, nil
}

//func (c *Client) GetClusters() (ClusterList, error) {
//	clusters := make(ClusterList, 0)
//	res, err := c.ListClusters(&eks.ListClustersInput{})
//	if err != nil {
//		return nil, err
//	}
//
//	for _, cc := range res.Clusters {
//		ngs, err := c.GetNodeGroups(*cc)
//		if err != nil {
//			return nil, err
//		}
//		if len(ngs) >= 1 {
//			clusters = append(clusters, &Cluster{
//				Name:       *cc,
//				NodeGroups: ngs,
//			})
//		}
//	}
//
//	return clusters, nil
//}
