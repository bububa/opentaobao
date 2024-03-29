package icbulogistics

import (
	"sync"
)

// PlaceOrderDto 结构体
type PlaceOrderDto struct {
	// 货品信息
	CargoList []CargoList `json:"cargo_list,omitempty" xml:"cargo_list>cargo_list,omitempty"`
	// 包裹信息
	PackageList []PackageList `json:"package_list,omitempty" xml:"package_list>package_list,omitempty"`
	// 起始地邮编
	OriginZipCode string `json:"origin_zip_code,omitempty" xml:"origin_zip_code,omitempty"`
	// 目的地国家
	DestinationCountryCode string `json:"destination_country_code,omitempty" xml:"destination_country_code,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 产品编码
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 目的地邮编
	DestinationZipCode string `json:"destination_zip_code,omitempty" xml:"destination_zip_code,omitempty"`
	// 发货批次ID
	SupplyChainBizId string `json:"supply_chain_biz_id,omitempty" xml:"supply_chain_biz_id,omitempty"`
	// 交易单号（例如阿里国际站的信保单ID），注意此字段不为空时，trade_platform字段必填（默认为ICBU）
	TradeBizId string `json:"trade_biz_id,omitempty" xml:"trade_biz_id,omitempty"`
	// 跨境电商平台代码：ICBU（阿里巴巴国际站）、ALIEXPRESS（速卖通）、AMAZON（亚马逊）等
	TradePlatform string `json:"trade_platform,omitempty" xml:"trade_platform,omitempty"`
	// 交货到仓快递信息
	DeliverWarehouseExpress *DeliverWarehouseExpressDto `json:"deliver_warehouse_express,omitempty" xml:"deliver_warehouse_express,omitempty"`
	// 发货人地址
	ConsignorAddress *ConsignorAddress `json:"consignor_address,omitempty" xml:"consignor_address,omitempty"`
	// 申报信息
	ExpressCustoms *ExpressCustomsDto `json:"express_customs,omitempty" xml:"express_customs,omitempty"`
	// 收货人地址
	ConsigneeAddress *ConsigneeAddress `json:"consignee_address,omitempty" xml:"consignee_address,omitempty"`
	// 备用字段（上门揽收地址），目前按发货人地址
	PickupAddress *ContactAddress `json:"pickup_address,omitempty" xml:"pickup_address,omitempty"`
	// 备用字段（退货地址），目前按退货申请指定地址
	ReturnAddress *ContactAddress `json:"return_address,omitempty" xml:"return_address,omitempty"`
	// 1
	NeedPickup bool `json:"need_pickup,omitempty" xml:"need_pickup,omitempty"`
}

var poolPlaceOrderDto = sync.Pool{
	New: func() any {
		return new(PlaceOrderDto)
	},
}

// GetPlaceOrderDto() 从对象池中获取PlaceOrderDto
func GetPlaceOrderDto() *PlaceOrderDto {
	return poolPlaceOrderDto.Get().(*PlaceOrderDto)
}

// ReleasePlaceOrderDto 释放PlaceOrderDto
func ReleasePlaceOrderDto(v *PlaceOrderDto) {
	v.CargoList = v.CargoList[:0]
	v.PackageList = v.PackageList[:0]
	v.OriginZipCode = ""
	v.DestinationCountryCode = ""
	v.WarehouseCode = ""
	v.ProductCode = ""
	v.DestinationZipCode = ""
	v.SupplyChainBizId = ""
	v.TradeBizId = ""
	v.TradePlatform = ""
	v.DeliverWarehouseExpress = nil
	v.ConsignorAddress = nil
	v.ExpressCustoms = nil
	v.ConsigneeAddress = nil
	v.PickupAddress = nil
	v.ReturnAddress = nil
	v.NeedPickup = false
	poolPlaceOrderDto.Put(v)
}
