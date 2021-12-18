package cloud_test

import (
	"testing"

	"github.com/devinturner/eks-tool/pkg/cloud"
)

func TestGetClusters(t *testing.T) {
	t.Run("us-east-1", func(t *testing.T) {
		c := cloud.NewClient("us-east-1")
		clusters, err := c.GetClusters()
		if err != nil {
			t.Error(err)
		} else {
			for _, cluster := range clusters {
				t.Log(cluster.Name)
			}
		}
	})
}
