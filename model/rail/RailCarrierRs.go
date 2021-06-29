package rail

// RailCarrierRS 
type RailCarrierRS struct {

    // 铁路运营公司Code
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 铁路运营公司Logo
    
    Logo   string `json:"logo,omitempty" xml:"logo,omitempty"`
    

    // 铁路运营公司中文名
    
    CnName   string `json:"cn_name,omitempty" xml:"cn_name,omitempty"`
    

    // 铁路运营公司英文名
    
    EnName   string `json:"en_name,omitempty" xml:"en_name,omitempty"`
    

    // 关联or归属铁路局编码
    
    RailWayCode   string `json:"rail_way_code,omitempty" xml:"rail_way_code,omitempty"`
    

}
