package mos

// OperatorUserInfo 结构体
type OperatorUserInfo struct {
	// 阿里工号
	AliWorkNo string `json:"ali_work_no,omitempty" xml:"ali_work_no,omitempty"`
	// 银泰工号
	CompWorkNo string `json:"comp_work_no,omitempty" xml:"comp_work_no,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// operator id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 淘宝id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
