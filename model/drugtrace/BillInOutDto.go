package drugtrace

// BillInOutDto 结构体
type BillInOutDto struct {
	// 收货单位名称
	ToUserName string `json:"to_user_name,omitempty" xml:"to_user_name,omitempty"`
	// 发货单位名称
	FromUserName string `json:"from_user_name,omitempty" xml:"from_user_name,omitempty"`
	// 单据编号
	BillCode string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
	// 入出库时间
	BillTimeStr string `json:"bill_time_str,omitempty" xml:"bill_time_str,omitempty"`
	// 单据类型
	BillTypeStr string `json:"bill_type_str,omitempty" xml:"bill_type_str,omitempty"`
	// 文件名
	UploadFileName string `json:"upload_file_name,omitempty" xml:"upload_file_name,omitempty"`
	// 创建时间
	CrtDate string `json:"crt_date,omitempty" xml:"crt_date,omitempty"`
	// 送达日期
	EtlTime string `json:"etl_time,omitempty" xml:"etl_time,omitempty"`
	// 是否匹配
	OrderIsMatched string `json:"order_is_matched,omitempty" xml:"order_is_matched,omitempty"`
	// 代理物流企业
	AgentUserName string `json:"agent_user_name,omitempty" xml:"agent_user_name,omitempty"`
	// 代理物流企业ID
	AgentRefUserId string `json:"agent_ref_user_id,omitempty" xml:"agent_ref_user_id,omitempty"`
	// 收货企业ID
	ToUserId string `json:"to_user_id,omitempty" xml:"to_user_id,omitempty"`
	// 发货企业ID
	FromUserId string `json:"from_user_id,omitempty" xml:"from_user_id,omitempty"`
	// 货主
	RefUserId string `json:"ref_user_id,omitempty" xml:"ref_user_id,omitempty"`
	// 药品标识
	DrugEntBaseInfoId string `json:"drug_ent_base_info_id,omitempty" xml:"drug_ent_base_info_id,omitempty"`
	// 单据处理日期
	ProcessDate string `json:"process_date,omitempty" xml:"process_date,omitempty"`
	// 单据时间
	BillTime string `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
	// 生产日期
	ProduceDate string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
	// 有效期
	ExprieDate string `json:"exprie_date,omitempty" xml:"exprie_date,omitempty"`
	// 制剂数量
	RemnantCount string `json:"remnant_count,omitempty" xml:"remnant_count,omitempty"`
	// 码数量
	CodeCount string `json:"code_count,omitempty" xml:"code_count,omitempty"`
	// produceBatchNo
	ProduceBatchNo string `json:"produce_batch_no,omitempty" xml:"produce_batch_no,omitempty"`
	// 旧单据id
	OldBillId string `json:"old_bill_id,omitempty" xml:"old_bill_id,omitempty"`
	// billType
	BillType string `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
	// 单据id
	BillId string `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
}
