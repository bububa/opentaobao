package train

import (
	"sync"
)

// StudentInfo 结构体
type StudentInfo struct {
	// demo
	Card string `json:"card,omitempty" xml:"card,omitempty"`
	// demo
	Classes string `json:"classes,omitempty" xml:"classes,omitempty"`
	// demo
	DepartMent string `json:"depart_ment,omitempty" xml:"depart_ment,omitempty"`
	// demo
	EductionalSystem string `json:"eductional_system,omitempty" xml:"eductional_system,omitempty"`
	// demo
	EntranceYear string `json:"entrance_year,omitempty" xml:"entrance_year,omitempty"`
	// demo
	FromCity string `json:"from_city,omitempty" xml:"from_city,omitempty"`
	// demo
	SchoolName string `json:"school_name,omitempty" xml:"school_name,omitempty"`
	// demo
	SchoolProvince string `json:"school_province,omitempty" xml:"school_province,omitempty"`
	// demo
	StudentNo string `json:"student_no,omitempty" xml:"student_no,omitempty"`
	// demo
	ToCity string `json:"to_city,omitempty" xml:"to_city,omitempty"`
}

var poolStudentInfo = sync.Pool{
	New: func() any {
		return new(StudentInfo)
	},
}

// GetStudentInfo() 从对象池中获取StudentInfo
func GetStudentInfo() *StudentInfo {
	return poolStudentInfo.Get().(*StudentInfo)
}

// ReleaseStudentInfo 释放StudentInfo
func ReleaseStudentInfo(v *StudentInfo) {
	v.Card = ""
	v.Classes = ""
	v.DepartMent = ""
	v.EductionalSystem = ""
	v.EntranceYear = ""
	v.FromCity = ""
	v.SchoolName = ""
	v.SchoolProvince = ""
	v.StudentNo = ""
	v.ToCity = ""
	poolStudentInfo.Put(v)
}
