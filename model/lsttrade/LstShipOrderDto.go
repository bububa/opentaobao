package lsttrade

import (
	"sync"
)

// LstShipOrderDto 结构体
type LstShipOrderDto struct {
	// 主发货单号
	MainShipOrderId string `json:"main_ship_order_id,omitempty" xml:"main_ship_order_id,omitempty"`
	// 交易订单号
	MainOrderId string `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 发货单是否作废
	Disused string `json:"disused,omitempty" xml:"disused,omitempty"`
	// 发货单来源
	SourceName string `json:"source_name,omitempty" xml:"source_name,omitempty"`
	// 作废后生成新主订单号
	NewMainOrderId string `json:"new_main_order_id,omitempty" xml:"new_main_order_id,omitempty"`
	// 子交易订单号
	SubOrderId string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 商品货号
	CargoNum string `json:"cargo_num,omitempty" xml:"cargo_num,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 单价
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 实付金额
	ItemAmount string `json:"item_amount,omitempty" xml:"item_amount,omitempty"`
	// 是否申请退款
	RefundInfo string `json:"refund_info,omitempty" xml:"refund_info,omitempty"`
	// 创建时间
	GmtCreateTime string `json:"gmt_create_time,omitempty" xml:"gmt_create_time,omitempty"`
	// 门店名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 门店编码或ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 门店别名
	ShopAliasName string `json:"shop_alias_name,omitempty" xml:"shop_alias_name,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 发货单状态
	StatusName string `json:"status_name,omitempty" xml:"status_name,omitempty"`
	// 发货时间
	LoadTime string `json:"load_time,omitempty" xml:"load_time,omitempty"`
	// 签收时间
	SignTime string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
	// 车辆信息
	VehicleInfo string `json:"vehicle_info,omitempty" xml:"vehicle_info,omitempty"`
	// 配送司机
	DriverInfo string `json:"driver_info,omitempty" xml:"driver_info,omitempty"`
	// 配送商名称
	DistributorName string `json:"distributor_name,omitempty" xml:"distributor_name,omitempty"`
	// 钉钉手机号
	DistributorPhone string `json:"distributor_phone,omitempty" xml:"distributor_phone,omitempty"`
	// 打印次数
	PrintTimes string `json:"print_times,omitempty" xml:"print_times,omitempty"`
	// 创建时间
	GmtUpdateTime string `json:"gmt_update_time,omitempty" xml:"gmt_update_time,omitempty"`
}

var poolLstShipOrderDto = sync.Pool{
	New: func() any {
		return new(LstShipOrderDto)
	},
}

// GetLstShipOrderDto() 从对象池中获取LstShipOrderDto
func GetLstShipOrderDto() *LstShipOrderDto {
	return poolLstShipOrderDto.Get().(*LstShipOrderDto)
}

// ReleaseLstShipOrderDto 释放LstShipOrderDto
func ReleaseLstShipOrderDto(v *LstShipOrderDto) {
	v.MainShipOrderId = ""
	v.MainOrderId = ""
	v.Disused = ""
	v.SourceName = ""
	v.NewMainOrderId = ""
	v.SubOrderId = ""
	v.ItemName = ""
	v.Barcode = ""
	v.CargoNum = ""
	v.Quantity = ""
	v.Price = ""
	v.ItemAmount = ""
	v.RefundInfo = ""
	v.GmtCreateTime = ""
	v.ShopName = ""
	v.ShopId = ""
	v.ShopAliasName = ""
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	v.DetailAddress = ""
	v.StatusName = ""
	v.LoadTime = ""
	v.SignTime = ""
	v.VehicleInfo = ""
	v.DriverInfo = ""
	v.DistributorName = ""
	v.DistributorPhone = ""
	v.PrintTimes = ""
	v.GmtUpdateTime = ""
	poolLstShipOrderDto.Put(v)
}
