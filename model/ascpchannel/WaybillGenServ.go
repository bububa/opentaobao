package ascpchannel

import (
	"sync"
)

// WaybillGenServ 结构体
type WaybillGenServ struct {
	// 服务类型：0为送装一体，1为只送到家不安装，2为只送不装到楼下，3为自提
	DeliveryType string `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
}

var poolWaybillGenServ = sync.Pool{
	New: func() any {
		return new(WaybillGenServ)
	},
}

// GetWaybillGenServ() 从对象池中获取WaybillGenServ
func GetWaybillGenServ() *WaybillGenServ {
	return poolWaybillGenServ.Get().(*WaybillGenServ)
}

// ReleaseWaybillGenServ 释放WaybillGenServ
func ReleaseWaybillGenServ(v *WaybillGenServ) {
	v.DeliveryType = ""
	poolWaybillGenServ.Put(v)
}
