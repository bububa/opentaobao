package drugtrace

// BillChkInOutDo 
type BillChkInOutDo struct {
    // 单据类型
    BillType   string `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
    // 单号号码
    BillCode   string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
    // 发货企业ID
    FromUserId   string `json:"from_user_id,omitempty" xml:"from_user_id,omitempty"`
    // 企业名称
    RefUserName   string `json:"ref_user_name,omitempty" xml:"ref_user_name,omitempty"`
    // 企业ID
    RefUserId   string `json:"ref_user_id,omitempty" xml:"ref_user_id,omitempty"`
    // 生产日期
    ProduceDate   string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
    // 上传文件名
    UploadFileName   string `json:"upload_file_name,omitempty" xml:"upload_file_name,omitempty"`
    // 发货单位
    FromUserName   string `json:"from_user_name,omitempty" xml:"from_user_name,omitempty"`
    // 收货单位
    ToUserId   string `json:"to_user_id,omitempty" xml:"to_user_id,omitempty"`
    // 生产企业ID
    ProduceEntId   string `json:"produce_ent_id,omitempty" xml:"produce_ent_id,omitempty"`
    // 单据时间
    BillTime   string `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
    // 角色类型
    UserRoleType   string `json:"user_role_type,omitempty" xml:"user_role_type,omitempty"`
    // 处理日期
    ProcessDate   string `json:"process_date,omitempty" xml:"process_date,omitempty"`
    // 单据ID
    BillId   string `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
    // 收货单位
    ToUserName   string `json:"to_user_name,omitempty" xml:"to_user_name,omitempty"`
    // 代理企业
    AgentUserName   string `json:"agent_user_name,omitempty" xml:"agent_user_name,omitempty"`
    // 代理企业ID
    AgentRefUserId   string `json:"agent_ref_user_id,omitempty" xml:"agent_ref_user_id,omitempty"`
    // 单据类型
    BillTypeName   string `json:"bill_type_name,omitempty" xml:"bill_type_name,omitempty"`
    // 收货单位ID
    ToRefUserId   string `json:"to_ref_user_id,omitempty" xml:"to_ref_user_id,omitempty"`
    // 发货单位ID
    FromRefUserId   string `json:"from_ref_user_id,omitempty" xml:"from_ref_user_id,omitempty"`
}
