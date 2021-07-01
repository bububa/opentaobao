package aliqin

// AlibabaAliqinFcIotQrycardModel 结构体
type AlibabaAliqinFcIotQrycardModel struct {
	// 流量类型(6700001:上网流量)
	ResourceType string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// 总流量(KB)
	FlowResource int64 `json:"flow_resource,omitempty" xml:"flow_resource,omitempty"`
	// 剩余流量(KB)
	RestOfFlow int64 `json:"rest_of_flow,omitempty" xml:"rest_of_flow,omitempty"`
	// 资源名称
	ResName string `json:"res_name,omitempty" xml:"res_name,omitempty"`
	// 生效时间
	ValidDate string `json:"valid_date,omitempty" xml:"valid_date,omitempty"`
	// 失效时间
	ExpireDate string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 已使用流量(KB)
	FlowUsed int64 `json:"flow_used,omitempty" xml:"flow_used,omitempty"`
}
