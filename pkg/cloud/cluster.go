package cloud

import "github.com/aws/aws-sdk-go/service/eks"

type Cluster struct {
	Name       string       `json:"name"`
	NodeGroups []*NodeGroup `json:"node_groups,omitempty"`
}

func (c *Client) GetClusters() ([]*Cluster, error) {
	clusters := make([]*Cluster, 0)
	res, err := c.ListClusters(&eks.ListClustersInput{})
	if err != nil {
		return nil, err
	}

	for _, cc := range res.Clusters {
		clusters = append(clusters, &Cluster{
			Name: *cc,
		})
	}
	return clusters, nil
}
