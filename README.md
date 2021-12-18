# eks-tool

A quick and dirty tool to list out all EKS clusters with active nodegroups in all regions in json format.

The scheme is as follows:

```json
{
  "regions": [
    {
      "name": "region",
      "clusters": [
        {
          "name": "cluster-1",
          "node_groups": [
            {
              "name": "ng-1",
              "node_count": 1
            }
          ]
        }
      ]
    }
  ]
}
```

This should allow a user to leverage `jq` or some other json processor to parse for valuable info. 

