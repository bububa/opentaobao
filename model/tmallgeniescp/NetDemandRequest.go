package tmallgeniescp

// NetDemandRequest 
type NetDemandRequest struct {

    // 扩展参数
    
    RequestExtendJson   string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
    

    // 对象
    
    NetDemandDTOs   []NetDemandDto `json:"net_demand_d_t_os,omitempty" xml:"net_demand_d_t_os,omitempty"`
    

}
