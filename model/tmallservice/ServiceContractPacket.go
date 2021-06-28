package tmallservice

// ServiceContractPacket 
/* model for simplify = false
type ServiceContractPacket struct {

    // 服务名字
    
    Name   string `json:"name,omitempty"`
    

    // 描述
    
    Desc   string `json:"desc,omitempty"`
    

    // 合同类服务列表
    
    ServiceList  struct {
        ServiceContractDO  []ServiceContractDO `json:"service_contract_do,omitempty"`
    } `json:"service_list,omitempty"`
    

}
*/

// ServiceContractPacket 
type ServiceContractPacket struct {

    // 服务名字
    Name   string `json:"name,omitempty"`

    // 描述
    Desc   string `json:"desc,omitempty"`

    // 合同类服务列表
    ServiceList   []ServiceContractDO `json:"service_list,omitempty"`

}
