package alidoc

// RxDoctorDto 结构体
type RxDoctorDto struct {
	// 医生部门
	DepartName string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 医生姓名
	DoctorName string `json:"doctor_name,omitempty" xml:"doctor_name,omitempty"`
}
