package drugtrace

// CodeActiveProcessDTO 
type CodeActiveProcessDTO struct {
    // 错误码信息
    NoteInfo   string `json:"note_info,omitempty" xml:"note_info,omitempty"`
    // 删除原因
    DeleteReason   string `json:"delete_reason,omitempty" xml:"delete_reason,omitempty"`
    // 操作人
    OperName   string `json:"oper_name,omitempty" xml:"oper_name,omitempty"`
    // 码
    PiatsCode   string `json:"piats_code,omitempty" xml:"piats_code,omitempty"`
    // 处理日期
    ProcessDate   string `json:"process_date,omitempty" xml:"process_date,omitempty"`
    // 制剂单位
    PrepnUnit   int64 `json:"prepn_unit,omitempty" xml:"prepn_unit,omitempty"`
    // 企业ID
    RefEntId   string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
    // 处理错误原因
    ProcessDes   string `json:"process_des,omitempty" xml:"process_des,omitempty"`
    // 关联关系文件名
    FileName   string `json:"file_name,omitempty" xml:"file_name,omitempty"`
    // 处理状态
    ProcessFlag   string `json:"process_flag,omitempty" xml:"process_flag,omitempty"`
    // 上传时间
    UploadDate   string `json:"upload_date,omitempty" xml:"upload_date,omitempty"`
    // 上传状态
    UploadFlag   string `json:"upload_flag,omitempty" xml:"upload_flag,omitempty"`
    // 有效期至
    ExprieDate   string `json:"exprie_date,omitempty" xml:"exprie_date,omitempty"`
    // 生产时间
    ProduceDate   string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
    // 最小赋码包装单位编码
    PkgUnitCode   string `json:"pkg_unit_code,omitempty" xml:"pkg_unit_code,omitempty"`
    // 小码数量
    SmallNum   string `json:"small_num,omitempty" xml:"small_num,omitempty"`
    // 最大包装数量
    OtherNum   string `json:"other_num,omitempty" xml:"other_num,omitempty"`
    // 激活数量
    ActiveCount   string `json:"active_count,omitempty" xml:"active_count,omitempty"`
    // 制剂规格
    PrepnSpec   string `json:"prepn_spec,omitempty" xml:"prepn_spec,omitempty"`
    // 包装比例
    PkgRatio   string `json:"pkg_ratio,omitempty" xml:"pkg_ratio,omitempty"`
    // 最小赋码包装单位描述
    PkgUnitDesc   string `json:"pkg_unit_desc,omitempty" xml:"pkg_unit_desc,omitempty"`
    // 包装规格
    PkgSpec   string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
    // 药品批号
    ProduceBatchNo   string `json:"produce_batch_no,omitempty" xml:"produce_batch_no,omitempty"`
    // 剂型描述
    PrepnTypeDesc   string `json:"prepn_type_desc,omitempty" xml:"prepn_type_desc,omitempty"`
    // 药品通用名
    PhysicName   string `json:"physic_name,omitempty" xml:"physic_name,omitempty"`
    // 药品信息
    PhysicInfo   string `json:"physic_info,omitempty" xml:"physic_info,omitempty"`
    // 生产企业标识
    ProdSeqNo   string `json:"prod_seq_no,omitempty" xml:"prod_seq_no,omitempty"`
    // 激活信息主键
    ActiveInfoSeqNo   string `json:"active_info_seq_no,omitempty" xml:"active_info_seq_no,omitempty"`
}
