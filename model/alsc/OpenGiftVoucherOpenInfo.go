package alsc

import (
	"sync"
)

// OpenGiftVoucherOpenInfo 结构体
type OpenGiftVoucherOpenInfo struct {
	// 券模板ID
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
}

var poolOpenGiftVoucherOpenInfo = sync.Pool{
	New: func() any {
		return new(OpenGiftVoucherOpenInfo)
	},
}

// GetOpenGiftVoucherOpenInfo() 从对象池中获取OpenGiftVoucherOpenInfo
func GetOpenGiftVoucherOpenInfo() *OpenGiftVoucherOpenInfo {
	return poolOpenGiftVoucherOpenInfo.Get().(*OpenGiftVoucherOpenInfo)
}

// ReleaseOpenGiftVoucherOpenInfo 释放OpenGiftVoucherOpenInfo
func ReleaseOpenGiftVoucherOpenInfo(v *OpenGiftVoucherOpenInfo) {
	v.VoucherId = ""
	v.Num = 0
	poolOpenGiftVoucherOpenInfo.Put(v)
}
