package tmallgeniescp

// NetDemandRawRequest 结构体
type NetDemandRawRequest struct {
	// 对象
	NetDemandDTOs []NetDemandRawDto `json:"net_demand_d_t_os,omitempty" xml:"net_demand_d_t_os>net_demand_raw_dto,omitempty"`
	// 扩展参数
	RequestExtendJson string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}
