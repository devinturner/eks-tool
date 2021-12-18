package cloud

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/eks"
)

type Region struct {
	Name     string     `json:"name,omitempty"`
	Clusters []*Cluster `json:"clusters,omitempty"`
}

func GetRegions() []*Region {
	out := make([]*Region, 0)
	partition := endpoints.AwsPartition()
	for rn, r := range partition.Regions() {
		for _, svc := range r.Services() {
			if strings.EqualFold(svc.ID(), eks.ServiceID) {
				out = append(out, &Region{Name: rn})
			}
		}
	}
	return out
}

func FilterRegions(regions []*Region, f func(*Region) bool) []*Region {
	out := make([]*Region, 0)
	for _, r := range regions {
		if f(r) {
			out = append(out, r)
		}
	}
	return out
}
