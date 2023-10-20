package alihealth2

// AlibabaalihealthtracecodesellerbillresultsearchResult 结构体
type AlibabaalihealthtracecodesellerbillresultsearchResult struct {
	// 单据编号
	BillCode string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
	// 单据处理时间
	GmtModified int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 单据处理状态
	TaskStatus int64 `json:"task_status,omitempty" xml:"task_status,omitempty"`
	// 单据上传时间
	GmtCreate int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 商家id
	OperEntId int64 `json:"oper_ent_id,omitempty" xml:"oper_ent_id,omitempty"`
}
