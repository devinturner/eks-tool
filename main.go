package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/devinturner/eks-tool/pkg/cloud"
	log "github.com/sirupsen/logrus"
)

var (
	verbose bool
)

func init() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {
	flag.BoolVar(&verbose, "verbose", false, "increase output verbosity")
	flag.Parse()

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	cl := new(cloud.Cloud)
	regions := cloud.FilterRegions(cloud.GetRegions(), func(r *cloud.Region) bool {
		for _, er := range []string{
			"eu-south-1",
			"af-south-1",
			"me-south-1",
			"ap-east-1",
		} {
			if r.Name == er {
				log.WithFields(log.Fields{
					"region": r.Name,
				}).Debug("excluding region from search")
				return false
			}
		}
		return true
	})
	for _, r := range regions {
		rc := cloud.NewClient(r.Name)

		log.WithFields(log.Fields{
			"region": r.Name,
		}).Debug("fetching clusters")

		clusters, err := rc.GetClusters()
		if err != nil {
			log.WithFields(log.Fields{
				"region": r.Name,
			}).Fatal(err)
		}
		for _, cluster := range clusters {
			log.WithFields(log.Fields{
				"region":  r.Name,
				"cluster": cluster.Name,
			}).Debug("fetching nodegroups")

			ngs, err := rc.GetNodeGroups(cluster.Name)
			if err != nil {
				log.WithFields(log.Fields{
					"region":  r.Name,
					"cluster": cluster.Name,
				}).Fatal(err)
			}
			for _, ng := range ngs {
				log.WithFields(log.Fields{
					"region":    r.Name,
					"cluster":   cluster.Name,
					"nodegroup": ng.Name,
				}).Debug("fetching nodegroup details")
				n, err := rc.GetNodeGroup(cluster.Name, ng.Name)
				if err != nil {
					log.WithFields(log.Fields{
						"region":    r.Name,
						"cluster":   cluster.Name,
						"nodegroup": ng.Name,
					}).Fatal(err)
				}
				ng.NodeCount = n.NodeCount
			}
			cluster.NodeGroups = ngs
		}
		r.Clusters = clusters
	}
	cl.Regions = cloud.FilterRegions(regions, func(r *cloud.Region) bool {
		return len(r.Clusters) > 0
	})

	b, err := json.Marshal(cl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
