package traveltrade

import (
	"sync"
)

// NormalVisaAppointmentInfo 结构体
type NormalVisaAppointmentInfo struct {
	// 必填，预约面试信pdf文件名称。具体的pdf文件字节流信息请设置到父级参数的 fileBytes字段！！！
	BookFileName string `json:"book_file_name,omitempty" xml:"book_file_name,omitempty"`
	// 必填，预约时间，格式：yyyy-MM-dd hh:mm:ss
	BookTime string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// 必填，预约地点
	BookPlace string `json:"book_place,omitempty" xml:"book_place,omitempty"`
}

var poolNormalVisaAppointmentInfo = sync.Pool{
	New: func() any {
		return new(NormalVisaAppointmentInfo)
	},
}

// GetNormalVisaAppointmentInfo() 从对象池中获取NormalVisaAppointmentInfo
func GetNormalVisaAppointmentInfo() *NormalVisaAppointmentInfo {
	return poolNormalVisaAppointmentInfo.Get().(*NormalVisaAppointmentInfo)
}

// ReleaseNormalVisaAppointmentInfo 释放NormalVisaAppointmentInfo
func ReleaseNormalVisaAppointmentInfo(v *NormalVisaAppointmentInfo) {
	v.BookFileName = ""
	v.BookTime = ""
	v.BookPlace = ""
	poolNormalVisaAppointmentInfo.Put(v)
}
