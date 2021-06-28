package tmallservice

// Value 
type Value struct {

    // 工单id列表
    
    WorkcardIdList   []int64 `json:"workcard_id_list,omitempty" xml:"workcard_id_list>int64,omitempty"`
    

    // 履约方式
    
    FulfilTypeCode   string `json:"fulfil_type_code,omitempty" xml:"fulfil_type_code,omitempty"`
    

    // 修改时间
    
    GmtModify   string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
    

    // 业务身份
    
    BizCode   string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
    

    // 下次联系时间
    
    GmtNextContact   string `json:"gmt_next_contact,omitempty" xml:"gmt_next_contact,omitempty"`
    

    // 备注
    
    Memo   string `json:"memo,omitempty" xml:"memo,omitempty"`
    

    // 预约信息
    
    Reservation   *Reservation `json:"reservation,omitempty" xml:"reservation,omitempty"`
    

    // 原因编码
    
    ReasonCode   int64 `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
    

    // 核销单id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 买家信息
    
    Buyer   *Buyer `json:"buyer,omitempty" xml:"buyer,omitempty"`
    

    // 服务提供者信息
    
    ServiceProvider   *ServiceProvider `json:"service_provider,omitempty" xml:"service_provider,omitempty"`
    

    // 核销单外部单号
    
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    

    // 原因描述
    
    ReasonDesc   string `json:"reason_desc,omitempty" xml:"reason_desc,omitempty"`
    

    // 状态编码
    
    StatusCode   string `json:"status_code,omitempty" xml:"status_code,omitempty"`
    

}
