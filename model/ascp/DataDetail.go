package ascp

import (
	"sync"
)

// DataDetail 结构体
type DataDetail struct {
	// 详情
	Detail []DetailItem `json:"detail,omitempty" xml:"detail>detail_item,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// ERP发货单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 订单发货地，省份
	SendProvince string `json:"send_province,omitempty" xml:"send_province,omitempty"`
	// 订单发货地，所在城市
	SendCity string `json:"send_city,omitempty" xml:"send_city,omitempty"`
	// 订单发货地，所在地区
	SendDistrict string `json:"send_district,omitempty" xml:"send_district,omitempty"`
	// 订单发货地，街道地址
	SendTown string `json:"send_town,omitempty" xml:"send_town,omitempty"`
	// 订单发货地地址编码（先识别编码，如果识别失败，解析地址）
	SendDivisionCode string `json:"send_division_code,omitempty" xml:"send_division_code,omitempty"`
	// 主交易单号
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 子交易单号
	SubTradeId string `json:"sub_trade_id,omitempty" xml:"sub_trade_id,omitempty"`
	// 订单类型 枚举： FENXIAO=分销订单 ，默认店铺零售订单。
	OrderFlag string `json:"order_flag,omitempty" xml:"order_flag,omitempty"`
	// 订单收货地，省份
	ReceiveProvince string `json:"receive_province,omitempty" xml:"receive_province,omitempty"`
	// 订单收货地，所在城市
	ReceiveCity string `json:"receive_city,omitempty" xml:"receive_city,omitempty"`
	// 订单收货地，所在地区
	ReceiveDistrict string `json:"receive_district,omitempty" xml:"receive_district,omitempty"`
	// 订单收货地，街道地址
	ReceiveTown string `json:"receive_town,omitempty" xml:"receive_town,omitempty"`
	// 订单收货地地址编码（先识别编码，如果识别失败，解析地址）
	ReceiveDivisionCode string `json:"receive_division_code,omitempty" xml:"receive_division_code,omitempty"`
	// 黑名单配品牌list,例：SF,YTO,STO
	BlackDeliveryCps string `json:"black_delivery_cps,omitempty" xml:"black_delivery_cps,omitempty"`
	// 白名单配品牌list,例：SF,YTO,STO
	WhiteDeliveryCps string `json:"white_delivery_cps,omitempty" xml:"white_delivery_cps,omitempty"`
	// ERP货品id
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 响应码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 响应信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 有SKU情况下必填；无SKU情况下为空（同一个itemid不可以同时既传有sku又传空的情况）
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 仓库编码，string（50)    卖家下唯一主键
	ErpWarehouseCode string `json:"erp_warehouse_code,omitempty" xml:"erp_warehouse_code,omitempty"`
}

var poolDataDetail = sync.Pool{
	New: func() any {
		return new(DataDetail)
	},
}

// GetDataDetail() 从对象池中获取DataDetail
func GetDataDetail() *DataDetail {
	return poolDataDetail.Get().(*DataDetail)
}

// ReleaseDataDetail 释放DataDetail
func ReleaseDataDetail(v *DataDetail) {
	v.Detail = v.Detail[:0]
	v.Result = ""
	v.OrderCode = ""
	v.SendProvince = ""
	v.SendCity = ""
	v.SendDistrict = ""
	v.SendTown = ""
	v.SendDivisionCode = ""
	v.TradeId = ""
	v.SubTradeId = ""
	v.OrderFlag = ""
	v.ReceiveProvince = ""
	v.ReceiveCity = ""
	v.ReceiveDistrict = ""
	v.ReceiveTown = ""
	v.ReceiveDivisionCode = ""
	v.BlackDeliveryCps = ""
	v.WhiteDeliveryCps = ""
	v.ScItemId = ""
	v.ErrCode = ""
	v.ErrMsg = ""
	v.ItemId = ""
	v.SkuId = ""
	v.ErpWarehouseCode = ""
	poolDataDetail.Put(v)
}
