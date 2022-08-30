package alidoc

// RxPatientDto 结构体
type RxPatientDto struct {
	// 身份证号
	IdCard string `json:"id_card,omitempty" xml:"id_card,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 年龄
	Age string `json:"age,omitempty" xml:"age,omitempty"`
	// 性别
	Sex string `json:"sex,omitempty" xml:"sex,omitempty"`
}
