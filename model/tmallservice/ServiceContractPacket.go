package tmallservice

// ServiceContractPacket 
type ServiceContractPacket struct {

    // 服务名字
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 描述
    
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    

    // 合同类服务列表
    
    ServiceList   []ServiceContractDO `json:"service_list,omitempty" xml:"service_list,omitempty"`
    

}
