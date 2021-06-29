package drugtrace

// FlowEntity 
type FlowEntity struct {

    // 发货企业名称
    
    FromEntName   string `json:"from_ent_name,omitempty" xml:"from_ent_name,omitempty"`
    

    // 企业所属省
    
    EntProvName   string `json:"ent_prov_name,omitempty" xml:"ent_prov_name,omitempty"`
    

    // 出库单单据号
    
    OutBillCode   string `json:"out_bill_code,omitempty" xml:"out_bill_code,omitempty"`
    

    // 入库单单据号
    
    InBillCode   string `json:"in_bill_code,omitempty" xml:"in_bill_code,omitempty"`
    

    // 出库单类型
    
    OutType   string `json:"out_type,omitempty" xml:"out_type,omitempty"`
    

    // 入库单类型
    
    InType   string `json:"in_type,omitempty" xml:"in_type,omitempty"`
    

    // 出库单id
    
    OutBillId   string `json:"out_bill_id,omitempty" xml:"out_bill_id,omitempty"`
    

    // 入库单id
    
    InBillId   string `json:"in_bill_id,omitempty" xml:"in_bill_id,omitempty"`
    

    // 出库时间
    
    OutDate   string `json:"out_date,omitempty" xml:"out_date,omitempty"`
    

    // 入库时间
    
    InDate   string `json:"in_date,omitempty" xml:"in_date,omitempty"`
    

    // 入库单类型编码
    
    InTypeCode   int64 `json:"in_type_code,omitempty" xml:"in_type_code,omitempty"`
    

    // 出库单类型编码
    
    OutTypeCode   string `json:"out_type_code,omitempty" xml:"out_type_code,omitempty"`
    

    // 收货企业id
    
    ToEntId   string `json:"to_ent_id,omitempty" xml:"to_ent_id,omitempty"`
    

    // 收货企业名称
    
    ToEntName   string `json:"to_ent_name,omitempty" xml:"to_ent_name,omitempty"`
    

    // 企业类型
    
    EntType   int64 `json:"ent_type,omitempty" xml:"ent_type,omitempty"`
    

    // 企业id
    
    EntId   string `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
    

    // 企业名称
    
    EntName   string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
    

    // 发货企业id
    
    FromEntId   string `json:"from_ent_id,omitempty" xml:"from_ent_id,omitempty"`
    

}
