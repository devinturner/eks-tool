package cloud_test

import (
	"testing"

	"github.com/devinturner/eks-tool/pkg/cloud"
)

func TestGetNodeGroups(t *testing.T) {
	cases := []struct {
		region  string
		cluster string
	}{
		{"us-east-1", "tarun-buildtest"},
		{"us-east-1", "sec-hamza-cluster"},
	}
	for _, tt := range cases {
		t.Run(tt.cluster, func(t *testing.T) {
			c := cloud.NewClient(tt.region)
			ngs, err := c.GetNodeGroups(tt.cluster)
			if err != nil {
				t.Error(err)
			}
			for _, ng := range ngs {
				t.Log(ng.Name)
			}
		})
	}
}

func TestGetNodeGroup(t *testing.T) {
	cases := []struct {
		region    string
		cluster   string
		nodegroup string
	}{
		{"us-east-1", "sec-hamza-cluster", "sec-hamza-pool-0"},
	}
	for _, tt := range cases {
		t.Run(tt.cluster, func(t *testing.T) {
			c := cloud.NewClient(tt.region)
			ng, err := c.GetNodeGroup(tt.cluster, tt.nodegroup)
			if err != nil {
				t.Error(err)
			}
			t.Log(ng.Name, ng.NodeCount)
		})
	}
}
