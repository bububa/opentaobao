package tmallgeniescp

// NetDemandRequest 结构体
type NetDemandRequest struct {
	// 对象
	NetDemandDTOs []NetDemandDto `json:"net_demand_d_t_os,omitempty" xml:"net_demand_d_t_os>net_demand_dto,omitempty"`
	// 扩展参数
	RequestExtendJson string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}
