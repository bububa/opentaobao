package alidoc

// RxDoctorTopDto 结构体
type RxDoctorTopDto struct {
	// 医生姓名
	DoctorName string `json:"doctor_name,omitempty" xml:"doctor_name,omitempty"`
	// 医生部门
	DepartName string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
}
