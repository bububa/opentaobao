package tmallhk

import (
	"sync"
)

// AwdcHrd 结构体
type AwdcHrd struct {
	// 参数extInfo
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 参数in
	In string `json:"in,omitempty" xml:"in,omitempty"`
	// 参数out
	Out string `json:"out,omitempty" xml:"out,omitempty"`
	// hrd证书信息
	ReportInfo string `json:"report_info,omitempty" xml:"report_info,omitempty"`
	// 参数reportNo
	ReportNo string `json:"report_no,omitempty" xml:"report_no,omitempty"`
}

var poolAwdcHrd = sync.Pool{
	New: func() any {
		return new(AwdcHrd)
	},
}

// GetAwdcHrd() 从对象池中获取AwdcHrd
func GetAwdcHrd() *AwdcHrd {
	return poolAwdcHrd.Get().(*AwdcHrd)
}

// ReleaseAwdcHrd 释放AwdcHrd
func ReleaseAwdcHrd(v *AwdcHrd) {
	v.ExtInfo = ""
	v.In = ""
	v.Out = ""
	v.ReportInfo = ""
	v.ReportNo = ""
	poolAwdcHrd.Put(v)
}
