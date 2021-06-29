package drugtrace

// AlibabaAlihealthDrugKytDestbillListModel 
type AlibabaAlihealthDrugKytDestbillListModel struct {

    // 药品类型
    
    DrugType   string `json:"drug_type,omitempty" xml:"drug_type,omitempty"`
    

    // 发货单位
    
    FromUserName   string `json:"from_user_name,omitempty" xml:"from_user_name,omitempty"`
    

    // 委托企业ref_ent_id
    
    AssRefEntId   string `json:"ass_ref_ent_id,omitempty" xml:"ass_ref_ent_id,omitempty"`
    

    // 直调收货企业
    
    DestEntName   string `json:"dest_ent_name,omitempty" xml:"dest_ent_name,omitempty"`
    

    // 委托企业名称
    
    AssEntName   string `json:"ass_ent_name,omitempty" xml:"ass_ent_name,omitempty"`
    

    // 配送企业ref_ent_id
    
    DisRefEntId   string `json:"dis_ref_ent_id,omitempty" xml:"dis_ref_ent_id,omitempty"`
    

    // 单据ID
    
    BillId   string `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
    

    // 委托企业ent_id
    
    AssEntId   string `json:"ass_ent_id,omitempty" xml:"ass_ent_id,omitempty"`
    

    // 配送企业ent_id
    
    DisEntId   string `json:"dis_ent_id,omitempty" xml:"dis_ent_id,omitempty"`
    

    // 收货企业名称
    
    ToUserName   string `json:"to_user_name,omitempty" xml:"to_user_name,omitempty"`
    

    // 配送企业名称
    
    DisEntName   string `json:"dis_ent_name,omitempty" xml:"dis_ent_name,omitempty"`
    

    // 审核状态，1：未审核；2：审核通过；3：审核失败
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 单据编号
    
    BillCode   string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
    

}
