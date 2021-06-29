package drugtrace

// DrugInfo 
type DrugInfo struct {
    // 有效期
    ExpiryDate   string `json:"expiry_date,omitempty" xml:"expiry_date,omitempty"`
    // 药品目录ID
    DrugBaseInfoId   string `json:"drug_base_info_id,omitempty" xml:"drug_base_info_id,omitempty"`
    // 生产日期
    ProductionDate   string `json:"production_date,omitempty" xml:"production_date,omitempty"`
    // 包装规格
    PkgSpec   string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
    // 药品ID
    DrugEntBaseInfoId   string `json:"drug_ent_base_info_id,omitempty" xml:"drug_ent_base_info_id,omitempty"`
    // 生产批号
    ProductionBatch   string `json:"production_batch,omitempty" xml:"production_batch,omitempty"`
    // 药品名称
    DrugName   string `json:"drug_name,omitempty" xml:"drug_name,omitempty"`
    // 批准文号
    LicenseNumber   string `json:"license_number,omitempty" xml:"license_number,omitempty"`
    // 制剂类型
    PrepnType   string `json:"prepn_type,omitempty" xml:"prepn_type,omitempty"`
    // 规格
    Specifications   string `json:"specifications,omitempty" xml:"specifications,omitempty"`
    // 生产企业
    Manufacturer   string `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
}
