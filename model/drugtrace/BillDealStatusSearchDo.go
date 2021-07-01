package drugtrace

// BillDealStatusSearchDo 
type BillDealStatusSearchDo struct {
    // 出入库号
    StoreInoutSeqNo   string `json:"store_inout_seq_no,omitempty" xml:"store_inout_seq_no,omitempty"`
    // 文件名
    ShortFileName   string `json:"short_file_name,omitempty" xml:"short_file_name,omitempty"`
    // 药品类型
    PhysicType   string `json:"physic_type,omitempty" xml:"physic_type,omitempty"`
    // 上传文件名
    UploadFileName   string `json:"upload_file_name,omitempty" xml:"upload_file_name,omitempty"`
    // 发货单位
    FromUserName   string `json:"from_user_name,omitempty" xml:"from_user_name,omitempty"`
    // 角色类型
    RoleType   string `json:"role_type,omitempty" xml:"role_type,omitempty"`
    // IC码
    IcCode   string `json:"ic_code,omitempty" xml:"ic_code,omitempty"`
    // 企业名称
    RefUserName   string `json:"ref_user_name,omitempty" xml:"ref_user_name,omitempty"`
    // 处理状态  0，处理中 1, 上传成功     3, 处理成功     4, 处理失败
    ResultType   string `json:"result_type,omitempty" xml:"result_type,omitempty"`
    // 上传标识
    UploadFlag   string `json:"upload_flag,omitempty" xml:"upload_flag,omitempty"`
    // 处理结果表状态(暂不用)
    ProcessFlag   string `json:"process_flag,omitempty" xml:"process_flag,omitempty"`
    // 单号号
    BillCode   string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
    // 单据类型
    BillType   string `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
    // 收货单位
    ToUserName   string `json:"to_user_name,omitempty" xml:"to_user_name,omitempty"`
    // 发货单位主键
    FromUserId   string `json:"from_user_id,omitempty" xml:"from_user_id,omitempty"`
    // 发货单位唯一标识
    FromRefUserId   string `json:"from_ref_user_id,omitempty" xml:"from_ref_user_id,omitempty"`
    // 收货单位主键
    ToUserId   string `json:"to_user_id,omitempty" xml:"to_user_id,omitempty"`
    // 用户唯一标识
    RefUserId   string `json:"ref_user_id,omitempty" xml:"ref_user_id,omitempty"`
    // 收货单位唯一标识
    ToRefUserId   string `json:"to_ref_user_id,omitempty" xml:"to_ref_user_id,omitempty"`
    // 用户主键
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 处理信息
    ProcessInfo   string `json:"process_info,omitempty" xml:"process_info,omitempty"`
    // 创建日期
    CrtDate   string `json:"crt_date,omitempty" xml:"crt_date,omitempty"`
    // 单据日期
    BillTime   string `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
    // 处理日期
    ProcessDate   string `json:"process_date,omitempty" xml:"process_date,omitempty"`
}
