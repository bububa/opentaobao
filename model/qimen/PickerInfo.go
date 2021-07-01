package qimen

// PickerInfo 结构体
type PickerInfo struct {
	// 公司名称
	Company string `json:"company,omitempty" xml:"company,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 固定电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 移动电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 证件号
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 车牌号
	CarNo string `json:"carNo,omitempty" xml:"carNo,omitempty"`
}
