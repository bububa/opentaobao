package alidoc

// RxPatientTopDto 
type RxPatientTopDto struct {
    // 性别
    Sex   string `json:"sex,omitempty" xml:"sex,omitempty"`
    // 年龄
    Age   int64 `json:"age,omitempty" xml:"age,omitempty"`
    // 姓名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 身份证号
    IdCard   string `json:"id_card,omitempty" xml:"id_card,omitempty"`
}
