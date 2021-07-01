package koubeimall

// ServiceInfoDto 结构体
type ServiceInfoDto struct {
	// 服务说明
	ServiceDesc string `json:"service_desc,omitempty" xml:"service_desc,omitempty"`
	// 门店服务tag
	ServiceTagList []string `json:"service_tag_list,omitempty" xml:"service_tag_list>string,omitempty"`
}
