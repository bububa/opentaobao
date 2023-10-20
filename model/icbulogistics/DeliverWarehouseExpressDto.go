package icbulogistics

import (
	"sync"
)

// DeliverWarehouseExpressDto 结构体
type DeliverWarehouseExpressDto struct {
	// 运单号
	TrackingNumbers []string `json:"tracking_numbers,omitempty" xml:"tracking_numbers>string,omitempty"`
	// 国内快递公司code
	LogisticsCompany string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
	// 包裹数量
	PackageQuantity string `json:"package_quantity,omitempty" xml:"package_quantity,omitempty"`
}

var poolDeliverWarehouseExpressDto = sync.Pool{
	New: func() any {
		return new(DeliverWarehouseExpressDto)
	},
}

// GetDeliverWarehouseExpressDto() 从对象池中获取DeliverWarehouseExpressDto
func GetDeliverWarehouseExpressDto() *DeliverWarehouseExpressDto {
	return poolDeliverWarehouseExpressDto.Get().(*DeliverWarehouseExpressDto)
}

// ReleaseDeliverWarehouseExpressDto 释放DeliverWarehouseExpressDto
func ReleaseDeliverWarehouseExpressDto(v *DeliverWarehouseExpressDto) {
	v.TrackingNumbers = v.TrackingNumbers[:0]
	v.LogisticsCompany = ""
	v.PackageQuantity = ""
	poolDeliverWarehouseExpressDto.Put(v)
}
