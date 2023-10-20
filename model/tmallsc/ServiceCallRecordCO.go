package tmallsc

// ServiceCallRecordCo 结构体
type ServiceCallRecordCo struct {
	// 修改日期
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 录音文件链接
	RecordLink string `json:"record_link,omitempty" xml:"record_link,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 呼叫记录的call_id
	CallId string `json:"call_id,omitempty" xml:"call_id,omitempty"`
	// 主叫拨打时间
	CallTime string `json:"call_time,omitempty" xml:"call_time,omitempty"`
	// 唯一标识
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
