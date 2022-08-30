package tmallservice

// ServiceContractPacket 结构体
type ServiceContractPacket struct {
	// 合同类服务列表
	ServiceList []ServiceContractDo `json:"service_list,omitempty" xml:"service_list>service_contract_do,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 服务名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
