package types

import (
	"time"

	dataprocpb "google.golang.org/genproto/googleapis/cloud/dataproc/v1"
)

type ClusterContainer struct {
	Clusters []*Cluster
}

func (c *ClusterContainer) Get(cloudType CloudType) []*Cluster {
	items := []*Cluster{}
	for _, item := range c.Clusters {
		if item.CloudType == cloudType {
			items = append(items, item)
		}
	}
	return items
}

func NewClusterContainer(clusters []*Cluster) *ClusterContainer {
	return &ClusterContainer{clusters}
}

// Cluster represents the Hadoop-based clusters on the cloud providers
type Cluster struct {
	Uuid      string                    `json:"ClusterUuid"`
	Name      string                    `json:"ClusterName"`
	Created   time.Time                 `json:"Created"`
	Tags      map[string]string         `json:"Tags"`
	Owner     string                    `json:"Owner"`
	CloudType CloudType                 `json:"CloudType"`
	State     State                     `json:"State"`
	Region    string                    `json:"Region"`
	Config    *dataprocpb.ClusterConfig `json:"ClusterConfig"`
}

// GetName returns the name of the cluster
func (cluster Cluster) GetName() string {
	return cluster.Name
}

// GetOwner returns the owner of the cluster
func (cluster Cluster) GetOwner() string {
	if cluster.Owner != "" {
		return cluster.Owner
	}
	if val, ok := cluster.Tags["Owner"]; ok {
		return val
	}
	if val, ok := cluster.Tags["owner"]; ok {
		return val
	}
	return ""
}

// GetCloudType returns the type of the cloud
func (cluster Cluster) GetCloudType() CloudType {
	return cluster.CloudType
}

// GetCreated returns the creation time of the cluster
func (cluster Cluster) GetCreated() time.Time {
	return cluster.Created
}

// GetItem returns the cluster struct itself
func (cluster Cluster) GetItem() interface{} {
	return cluster
}

// GetType returns the cluster's string representation
func (cluster Cluster) GetType() string {
	return "cluster"
}

// GetTags returns the cluster's tags
func (cluster Cluster) GetTags() Tags {
	return cluster.Tags
}
