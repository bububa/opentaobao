package drugtrace

// CodeRelationDetailInfoDo 
type CodeRelationDetailInfoDo struct {
    // 处理日期
    ProcessDate   string `json:"process_date,omitempty" xml:"process_date,omitempty"`
    // 关联模式[2 后关联]
    RelationType   string `json:"relation_type,omitempty" xml:"relation_type,omitempty"`
    // 备注
    Note   string `json:"note,omitempty" xml:"note,omitempty"`
    // 处理状态[1.待处理  3.处理成功 4.处理失败 ]
    ProcessFlag   string `json:"process_flag,omitempty" xml:"process_flag,omitempty"`
    // 上传者证书号
    UserCert   string `json:"user_cert,omitempty" xml:"user_cert,omitempty"`
    // 操作者姓名
    OperIcName   string `json:"oper_ic_name,omitempty" xml:"oper_ic_name,omitempty"`
    // 操作者编号
    OperIcCode   string `json:"oper_ic_code,omitempty" xml:"oper_ic_code,omitempty"`
    // 关联关系文件名
    FileName   string `json:"file_name,omitempty" xml:"file_name,omitempty"`
    // 上传状态
    UploadFlag   string `json:"upload_flag,omitempty" xml:"upload_flag,omitempty"`
    // 上传日期
    CrtDate   string `json:"crt_date,omitempty" xml:"crt_date,omitempty"`
    // 建立日期
    ActiveDate   string `json:"active_date,omitempty" xml:"active_date,omitempty"`
    // 最大包装数量
    OtherNum   string `json:"other_num,omitempty" xml:"other_num,omitempty"`
    // 小码数量
    SmallNum   string `json:"small_num,omitempty" xml:"small_num,omitempty"`
    // 激活数量(激活总量)
    ActiveCount   string `json:"active_count,omitempty" xml:"active_count,omitempty"`
    // 激活方式
    ActiveMethod   string `json:"active_method,omitempty" xml:"active_method,omitempty"`
    // 追溯码申请单
    ApplySeqNo   string `json:"apply_seq_no,omitempty" xml:"apply_seq_no,omitempty"`
    // 码激活文件上传信息标识
    CodeActiveInfoId   string `json:"code_active_info_id,omitempty" xml:"code_active_info_id,omitempty"`
    // 企业id
    EntId   string `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
}
