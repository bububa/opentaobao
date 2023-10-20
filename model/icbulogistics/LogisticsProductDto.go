package icbulogistics

import (
	"sync"
)

// LogisticsProductDto 结构体
type LogisticsProductDto struct {
	// 仓库名称
	WarehouseName string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 产品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 产品编码
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 时效类型
	DeliveryType string `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
	// 是否上门揽收
	Pickup bool `json:"pickup,omitempty" xml:"pickup,omitempty"`
}

var poolLogisticsProductDto = sync.Pool{
	New: func() any {
		return new(LogisticsProductDto)
	},
}

// GetLogisticsProductDto() 从对象池中获取LogisticsProductDto
func GetLogisticsProductDto() *LogisticsProductDto {
	return poolLogisticsProductDto.Get().(*LogisticsProductDto)
}

// ReleaseLogisticsProductDto 释放LogisticsProductDto
func ReleaseLogisticsProductDto(v *LogisticsProductDto) {
	v.WarehouseName = ""
	v.WarehouseCode = ""
	v.ProductName = ""
	v.ProductCode = ""
	v.DeliveryType = ""
	v.Pickup = false
	poolLogisticsProductDto.Put(v)
}
