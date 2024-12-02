package model

type AKS struct {
	Name           string `json:"name,omitempty" yaml:"name,omitempty"`
	Version        string `json:"version,omitempty" yaml:"version,omitempty"`
	Location       string `json:"location,omitempty" yaml:"location,omitempty"`
	DNSPrefix      string `json:"dns_prefix,omitempty" yaml:"dns_prefix,omitempty"`
	ResourceGroup  string `json:"resource_group,omitempty" yaml:"resource_group,omitempty"`
	SubscriptionID string `json:"subscription_id,omitempty" yaml:"subscription_id,omitempty"`
	PodCIDR        string `json:"pod_cidr,omitempty" yaml:"pod_cidr,omitempty"`
	ServiceCIDR    string `json:"service_cidr,omitempty" yaml:"service_cidr,omitempty"`
}

type ACA struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

type AzureNodepool struct {
	Name   string `json:"name,omitempty" yaml:"name,omitempty"`
	Min    int32  `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Max    int32  `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Count  int32  `json:"count,omitempty" yaml:"count,omitempty"`
	VMType string `json:"vm_type,omitempty" yaml:"vm_type,omitempty"`
	OSType string `json:"os_type,omitempty" yaml:"os_type,omitempty"`
}

type EKS struct {
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
	Region  string `json:"region,omitempty" yaml:"region,omitempty"`
}

type AWSManagedNodegroup struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

type GKE struct {
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
	Region  string `json:"region,omitempty" yaml:"region,omitempty"`
}

type GKENodepool struct {
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
