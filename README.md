# eks-tool

A quick and dirty tool to list out all EKS clusters with active nodegroups in all regions in json format.

The scheme is as follows:

```json
[
    {
        "name": "region",
        "clusters": [
            {
                "name": "cluster-1",
                "node_groups": [
                    "ng-1"
                ]
            }
        ]
    }
]
```

This should allow a user to leverage `jq` or some other json processor to parse for valuable info. For example, if a user is looking for all clusters in "us-west-2": 
```
go run cmd/main.go | jq '.[] | select( .name | contains("us-west-2"))'
```