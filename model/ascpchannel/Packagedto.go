package ascpchannel

import (
	"sync"
)

// Packagedto 结构体
type Packagedto struct {
	// 发货仓code
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 外部订单号
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 订单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 发货LP单号
	ConsignLgOrderCode string `json:"consign_lg_order_code,omitempty" xml:"consign_lg_order_code,omitempty"`
	// 配送公司code
	TmsResCode string `json:"tms_res_code,omitempty" xml:"tms_res_code,omitempty"`
	// 包裹号
	PackageCode string `json:"package_code,omitempty" xml:"package_code,omitempty"`
	// 包裹状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolPackagedto = sync.Pool{
	New: func() any {
		return new(Packagedto)
	},
}

// GetPackagedto() 从对象池中获取Packagedto
func GetPackagedto() *Packagedto {
	return poolPackagedto.Get().(*Packagedto)
}

// ReleasePackagedto 释放Packagedto
func ReleasePackagedto(v *Packagedto) {
	v.StoreCode = ""
	v.OutBizId = ""
	v.OrderCode = ""
	v.ConsignLgOrderCode = ""
	v.TmsResCode = ""
	v.PackageCode = ""
	v.Status = ""
	poolPackagedto.Put(v)
}
