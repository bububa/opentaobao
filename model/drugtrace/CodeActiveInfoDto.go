package drugtrace

// CodeActiveInfoDto 结构体
type CodeActiveInfoDto struct {
	// 处理标志
	ProcessFlag string `json:"process_flag,omitempty" xml:"process_flag,omitempty"`
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 关联关系类型
	RelationType string `json:"relation_type,omitempty" xml:"relation_type,omitempty"`
	// 紧急人
	UserCert string `json:"user_cert,omitempty" xml:"user_cert,omitempty"`
	// 生产编号
	ProdCode string `json:"prod_code,omitempty" xml:"prod_code,omitempty"`
	// 操作人姓名
	OperIcName string `json:"oper_ic_name,omitempty" xml:"oper_ic_name,omitempty"`
	// 上传文件名
	UploadFileName string `json:"upload_file_name,omitempty" xml:"upload_file_name,omitempty"`
	// 上传文件路径
	UploadFilePath string `json:"upload_file_path,omitempty" xml:"upload_file_path,omitempty"`
	// 激活时间
	ActiveDate string `json:"active_date,omitempty" xml:"active_date,omitempty"`
	// 企业ID
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 旧企业ID
	EntId string `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
	// 处理日期
	ProcessDate string `json:"process_date,omitempty" xml:"process_date,omitempty"`
	// 处理结束时间
	ProcessEndDate string `json:"process_end_date,omitempty" xml:"process_end_date,omitempty"`
	// 操作人编码
	OperIcCode string `json:"oper_ic_code,omitempty" xml:"oper_ic_code,omitempty"`
	// 上传标识
	UploadFlag string `json:"upload_flag,omitempty" xml:"upload_flag,omitempty"`
	// 激活时间
	CrtDate string `json:"crt_date,omitempty" xml:"crt_date,omitempty"`
	// 处理数量
	ProcessCount string `json:"process_count,omitempty" xml:"process_count,omitempty"`
	// 总激活数量
	ActiveCount int64 `json:"active_count,omitempty" xml:"active_count,omitempty"`
	// 最大包装数量
	OtherNum int64 `json:"other_num,omitempty" xml:"other_num,omitempty"`
	// 小码数量
	SmallNum int64 `json:"small_num,omitempty" xml:"small_num,omitempty"`
	// 关联关系文件上传日期
	CrtDateString string `json:"crt_date_string,omitempty" xml:"crt_date_string,omitempty"`
	// 单据id
	BillInId string `json:"bill_in_id,omitempty" xml:"bill_in_id,omitempty"`
	// 激活信息id
	CodeActiveInfoId string `json:"code_active_info_id,omitempty" xml:"code_active_info_id,omitempty"`
}
