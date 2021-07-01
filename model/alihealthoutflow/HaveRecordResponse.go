package alihealthoutflow

// HaveRecordResponse 结构体
type HaveRecordResponse struct {
	// 档案修改时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 档案创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 是否有效
	ValidStatus bool `json:"valid_status,omitempty" xml:"valid_status,omitempty"`
	// 是否存在
	ExistStatus bool `json:"exist_status,omitempty" xml:"exist_status,omitempty"`
}
