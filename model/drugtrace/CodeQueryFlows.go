package drugtrace

// CodeQueryFlows 
type CodeQueryFlows struct {
    // 入库单委托单位
    FromAssRefEntname   string `json:"from_ass_ref_entname,omitempty" xml:"from_ass_ref_entname,omitempty"`
    // 出库日期
    OutDate   string `json:"out_date,omitempty" xml:"out_date,omitempty"`
    // 企业名称
    EntName   string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
    // 出库单委托单位
    ToAssRefEntName   string `json:"to_ass_ref_ent_name,omitempty" xml:"to_ass_ref_ent_name,omitempty"`
    // 级别，0：生产，1：一级经销商，2：二级经销商，3：三级经销商
    Level   string `json:"level,omitempty" xml:"level,omitempty"`
    // 入库日期
    InDate   string `json:"in_date,omitempty" xml:"in_date,omitempty"`
    // 出库单类型
    OutType   string `json:"out_type,omitempty" xml:"out_type,omitempty"`
    // 出库单委托企业ID
    ToAssRefEntId   string `json:"to_ass_ref_ent_id,omitempty" xml:"to_ass_ref_ent_id,omitempty"`
    // 企业类型
    EntType   string `json:"ent_type,omitempty" xml:"ent_type,omitempty"`
    // 是否授权
    Authority   string `json:"authority,omitempty" xml:"authority,omitempty"`
    // 企业ID
    RefEntId   string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
    // 企业类型code
    EntTypeCode   string `json:"ent_type_code,omitempty" xml:"ent_type_code,omitempty"`
    // 行政区域
    Region   string `json:"region,omitempty" xml:"region,omitempty"`
    // 入库类型
    InType   string `json:"in_type,omitempty" xml:"in_type,omitempty"`
    // 入库委托企业ID
    FromAssRefEntId   string `json:"from_ass_ref_ent_id,omitempty" xml:"from_ass_ref_ent_id,omitempty"`
}
