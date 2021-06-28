package alicom

// OfflineOrderRecordExtInfo 
type OfflineOrderRecordExtInfo struct {

    // 网点地址
    
    AgentInfoAdress   string `json:"agent_info_adress,omitempty" xml:"agent_info_adress,omitempty"`
    

    // 网点名称
    
    AgentInfoName   string `json:"agent_info_name,omitempty" xml:"agent_info_name,omitempty"`
    

    // 网点id
    
    AgentInfoId   int64 `json:"agent_info_id,omitempty" xml:"agent_info_id,omitempty"`
    

    // 客户经理联系电话
    
    PhoneNo   string `json:"phone_no,omitempty" xml:"phone_no,omitempty"`
    

    // 客户经理姓名
    
    AgentName   string `json:"agent_name,omitempty" xml:"agent_name,omitempty"`
    

    // 客户经理id
    
    AgentId   int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
    

    // 省份
    
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    

    // 地市
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

}
