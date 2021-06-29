package drugtrace

// BillUpOutDetailDo 
type BillUpOutDetailDo struct {
    // 发货单位
    FromEntName   string `json:"from_ent_name,omitempty" xml:"from_ent_name,omitempty"`
    // 最小码量
    CodeCount   int64 `json:"code_count,omitempty" xml:"code_count,omitempty"`
    // 失效日期
    ExprieDate   string `json:"exprie_date,omitempty" xml:"exprie_date,omitempty"`
    // 厂商
    ProduceEntName   string `json:"produce_ent_name,omitempty" xml:"produce_ent_name,omitempty"`
    // 生产日期
    ProduceDate   string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
    // 生产批号
    ProduceBatchNo   string `json:"produce_batch_no,omitempty" xml:"produce_batch_no,omitempty"`
    // 包装规格
    PkgSpec   string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
    // 药品信息
    PhysicInfo   string `json:"physic_info,omitempty" xml:"physic_info,omitempty"`
    // 药品名称
    PhysicName   string `json:"physic_name,omitempty" xml:"physic_name,omitempty"`
    // 制剂数量
    PrepnCount   int64 `json:"prepn_count,omitempty" xml:"prepn_count,omitempty"`
    // 发货单位REF_ENT_ID
    FromRefUserId   string `json:"from_ref_user_id,omitempty" xml:"from_ref_user_id,omitempty"`
    // 收货单位REF_ENT_ID
    ToRefUserId   string `json:"to_ref_user_id,omitempty" xml:"to_ref_user_id,omitempty"`
    // 单据时间
    BillTime   string `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
    // 单据码
    BillCode   string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
    // 单据类型
    BillType   string `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
    // 发货企业
    ToUserName   string `json:"to_user_name,omitempty" xml:"to_user_name,omitempty"`
    // 收货企业
    FromUserName   string `json:"from_user_name,omitempty" xml:"from_user_name,omitempty"`
    // 失效日期格式化
    ExprieDateFormat   string `json:"exprie_date_format,omitempty" xml:"exprie_date_format,omitempty"`
    // 单据时间格式化
    BillTimeFormat   string `json:"bill_time_format,omitempty" xml:"bill_time_format,omitempty"`
    // 单据ID
    BillOutId   int64 `json:"bill_out_id,omitempty" xml:"bill_out_id,omitempty"`
    // 制剂单位
    PrepnUnit   int64 `json:"prepn_unit,omitempty" xml:"prepn_unit,omitempty"`
    // 制剂规格
    PrepnSpec   string `json:"prepn_spec,omitempty" xml:"prepn_spec,omitempty"`
    // 药品ID
    DrugEntBaseInfoId   string `json:"drug_ent_base_info_id,omitempty" xml:"drug_ent_base_info_id,omitempty"`
    // 生产日期格式化
    ProduceDateFormat   string `json:"produce_date_format,omitempty" xml:"produce_date_format,omitempty"`
    // 确认状态1未确认2已确认
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 收货企业ent_id
    ToUserId   string `json:"to_user_id,omitempty" xml:"to_user_id,omitempty"`
    // 发货企业ent_id
    FromUserId   string `json:"from_user_id,omitempty" xml:"from_user_id,omitempty"`
}
