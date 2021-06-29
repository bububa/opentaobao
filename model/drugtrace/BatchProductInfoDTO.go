package drugtrace

// BatchProductInfoDTO 
type BatchProductInfoDTO struct {
    // 生产企业名称
    EntName   string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
    // 包装规格
    PkgSpec   string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
    // 生产日期
    ProductDate   string `json:"product_date,omitempty" xml:"product_date,omitempty"`
    // 生产批号
    ProduceBatchNo   string `json:"produce_batch_no,omitempty" xml:"produce_batch_no,omitempty"`
    // 药品唯一ID
    DrugEntBaseId   string `json:"drug_ent_base_id,omitempty" xml:"drug_ent_base_id,omitempty"`
    // 药品通用名
    PhysicName   string `json:"physic_name,omitempty" xml:"physic_name,omitempty"`
    // 失效日期
    ExpireDate   string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
    // 生产企业唯一ID
    RefEntId   string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
    // 生产企业流水ID
    CurrEntId   string `json:"curr_ent_id,omitempty" xml:"curr_ent_id,omitempty"`
    // 剂型
    PrepnType   string `json:"prepn_type,omitempty" xml:"prepn_type,omitempty"`
    // 药品本位码
    SdcCode   string `json:"sdc_code,omitempty" xml:"sdc_code,omitempty"`
    // 制剂规格
    PrepnSpec   string `json:"prepn_spec,omitempty" xml:"prepn_spec,omitempty"`
    // 批准文号
    ApprovalLicenceNo   string `json:"approval_licence_no,omitempty" xml:"approval_licence_no,omitempty"`
}
