package drugtrace

// DrugEntBaseDto 
type DrugEntBaseDto struct {
    // 药品信息id
    DrugEntBaseId   string `json:"drug_ent_base_id,omitempty" xml:"drug_ent_base_id,omitempty"`
    // 药品基本信息id
    DrugEntBaseInfoId   string `json:"drug_ent_base_info_id,omitempty" xml:"drug_ent_base_info_id,omitempty"`
    // 批准文号
    ApprovalLicenceNo   string `json:"approval_licence_no,omitempty" xml:"approval_licence_no,omitempty"`
    // 药品名称
    PhysicName   string `json:"physic_name,omitempty" xml:"physic_name,omitempty"`
    // 药品类型描述
    PhysicTypeDesc   string `json:"physic_type_desc,omitempty" xml:"physic_type_desc,omitempty"`
    // 包装规格
    PkgSpecCrit   string `json:"pkg_spec_crit,omitempty" xml:"pkg_spec_crit,omitempty"`
    // 制剂规格
    PrepnSpec   string `json:"prepn_spec,omitempty" xml:"prepn_spec,omitempty"`
    // 剂型描述
    PrepnTypeDesc   string `json:"prepn_type_desc,omitempty" xml:"prepn_type_desc,omitempty"`
    // 有效期长度
    Exprie   string `json:"exprie,omitempty" xml:"exprie,omitempty"`
    // 商品名称
    ProdName   string `json:"prod_name,omitempty" xml:"prod_name,omitempty"`
    // 包装数量
    PkgNum   int64 `json:"pkg_num,omitempty" xml:"pkg_num,omitempty"`
    // 批准文号
    SdcCode   string `json:"sdc_code,omitempty" xml:"sdc_code,omitempty"`
    // 制剂单位类型描述
    PrepnUnitDesc   string `json:"prepn_unit_desc,omitempty" xml:"prepn_unit_desc,omitempty"`
    // 药品子类编码
    ProdCode   string `json:"prod_code,omitempty" xml:"prod_code,omitempty"`
    // 药品编码
    PhysicCode   string `json:"physic_code,omitempty" xml:"physic_code,omitempty"`
    // 药品详细类型
    PhysicDetailType   int64 `json:"physic_detail_type,omitempty" xml:"physic_detail_type,omitempty"`
    // 是否混批
    CodeActiveProductFlagStr   string `json:"code_active_product_flag_str,omitempty" xml:"code_active_product_flag_str,omitempty"`
    // 企业名称
    EntName   string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
    // 修改时间
    ModDate   string `json:"mod_date,omitempty" xml:"mod_date,omitempty"`
    // 创建时间
    CrtDate   string `json:"crt_date,omitempty" xml:"crt_date,omitempty"`
    // 是否授权0未授权,1已授权
    AuthorizerFlag   string `json:"authorizer_flag,omitempty" xml:"authorizer_flag,omitempty"`
    // 修改操作证书号
    ModIcCode   string `json:"mod_ic_code,omitempty" xml:"mod_ic_code,omitempty"`
    // 药品信息
    PhysicInfo   string `json:"physic_info,omitempty" xml:"physic_info,omitempty"`
    // 69码
    Ean13Code   string `json:"ean13_code,omitempty" xml:"ean13_code,omitempty"`
    // 包装规格
    PkgSpec   string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
    // 包装单位
    PkgUnit   int64 `json:"pkg_unit,omitempty" xml:"pkg_unit,omitempty"`
    // 制剂单位类型(详见码表)
    PrepnUnit   int64 `json:"prepn_unit,omitempty" xml:"prepn_unit,omitempty"`
    // 药品类型(详见码表) 1：特殊药品原料药，2：特殊药品制剂，3：普通药品，9：未分类
    PhysicType   int64 `json:"physic_type,omitempty" xml:"physic_type,omitempty"`
    // 剂型(制剂类型)(详见码表)
    PrepnType   int64 `json:"prepn_type,omitempty" xml:"prepn_type,omitempty"`
    // 申请码标识:1未申请,2已申请
    PhysicNatCode   int64 `json:"physic_nat_code,omitempty" xml:"physic_nat_code,omitempty"`
    // 状态 1正常,0停用
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 有效期单位(详见码表) 1:日，2：月，3：年
    ExprieUnit   int64 `json:"exprie_unit,omitempty" xml:"exprie_unit,omitempty"`
    // 药品有效期数(月/年)
    ExprieLife   int64 `json:"exprie_life,omitempty" xml:"exprie_life,omitempty"`
    // 年赋码量
    AnnCodeAmt   int64 `json:"ann_code_amt,omitempty" xml:"ann_code_amt,omitempty"`
    // 批准文号有效期至
    ApprovalLicenceExpiry   string `json:"approval_licence_expiry,omitempty" xml:"approval_licence_expiry,omitempty"`
    // 批准文号发放日期
    ApprovalLicenceDate   string `json:"approval_licence_date,omitempty" xml:"approval_licence_date,omitempty"`
    // 批准文号类型
    ApprovalLicenceType   int64 `json:"approval_licence_type,omitempty" xml:"approval_licence_type,omitempty"`
    // 药品批准文号ID(药品目录)
    DrugApprovalInfoId   string `json:"drug_approval_info_id,omitempty" xml:"drug_approval_info_id,omitempty"`
    // 药品基本信息ID
    DrugBaseInfoId   string `json:"drug_base_info_id,omitempty" xml:"drug_base_info_id,omitempty"`
    // 企业唯一标识
    RefEntId   string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
    // 企业主键id
    EntId   string `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
}
