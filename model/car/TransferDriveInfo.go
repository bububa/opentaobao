package car

import (
	"sync"
)

// TransferDriveInfo 结构体
type TransferDriveInfo struct {
	// 车牌号
	License string `json:"license,omitempty" xml:"license,omitempty"`
	// 司机姓名
	DriverName string `json:"driver_name,omitempty" xml:"driver_name,omitempty"`
	// 司机联系方式 真实号
	DriverPhone string `json:"driver_phone,omitempty" xml:"driver_phone,omitempty"`
	// 司机推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 行李信息
	Luggage string `json:"luggage,omitempty" xml:"luggage,omitempty"`
}

var poolTransferDriveInfo = sync.Pool{
	New: func() any {
		return new(TransferDriveInfo)
	},
}

// GetTransferDriveInfo() 从对象池中获取TransferDriveInfo
func GetTransferDriveInfo() *TransferDriveInfo {
	return poolTransferDriveInfo.Get().(*TransferDriveInfo)
}

// ReleaseTransferDriveInfo 释放TransferDriveInfo
func ReleaseTransferDriveInfo(v *TransferDriveInfo) {
	v.License = ""
	v.DriverName = ""
	v.DriverPhone = ""
	v.PushTime = ""
	v.Luggage = ""
	poolTransferDriveInfo.Put(v)
}
