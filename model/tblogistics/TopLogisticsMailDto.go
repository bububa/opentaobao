package tblogistics

// TopLogisticsMailDto 结构体
type TopLogisticsMailDto struct {
	// 物流节点列表
	TraceList []TopLogisticsNodeDto `json:"trace_list,omitempty" xml:"trace_list>top_logistics_node_dto,omitempty"`
	// 运单号
	OutSid string `json:"out_sid,omitempty" xml:"out_sid,omitempty"`
	// 物流公司
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 当前最新节点
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 交易单号
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}
