package tbtrade

import (
	"sync"
)

// ShipInfo 结构体
type ShipInfo struct {
	// 发货方式（小写）
	TransportType string `json:"transport_type,omitempty" xml:"transport_type,omitempty"`
}

var poolShipInfo = sync.Pool{
	New: func() any {
		return new(ShipInfo)
	},
}

// GetShipInfo() 从对象池中获取ShipInfo
func GetShipInfo() *ShipInfo {
	return poolShipInfo.Get().(*ShipInfo)
}

// ReleaseShipInfo 释放ShipInfo
func ReleaseShipInfo(v *ShipInfo) {
	v.TransportType = ""
	poolShipInfo.Put(v)
}
