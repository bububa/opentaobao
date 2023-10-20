package ascp

import (
	"sync"
)

// CourierInfos 结构体
type CourierInfos struct {
	// 网点小件员名称
	CourierName string `json:"courier_name,omitempty" xml:"courier_name,omitempty"`
	// 网点小件员编码（如工号等，服务商下唯一身份id）
	CourierNo string `json:"courier_no,omitempty" xml:"courier_no,omitempty"`
	// 小件员手机号
	CourierMobile string `json:"courier_mobile,omitempty" xml:"courier_mobile,omitempty"`
}

var poolCourierInfos = sync.Pool{
	New: func() any {
		return new(CourierInfos)
	},
}

// GetCourierInfos() 从对象池中获取CourierInfos
func GetCourierInfos() *CourierInfos {
	return poolCourierInfos.Get().(*CourierInfos)
}

// ReleaseCourierInfos 释放CourierInfos
func ReleaseCourierInfos(v *CourierInfos) {
	v.CourierName = ""
	v.CourierNo = ""
	v.CourierMobile = ""
	poolCourierInfos.Put(v)
}
