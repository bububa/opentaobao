package tbtrade

import (
	"sync"
)

// ServiceOrder 结构体
type ServiceOrder struct {
	// 服务详情的URL地址
	ServiceDetailUrl string `json:"service_detail_url,omitempty" xml:"service_detail_url,omitempty"`
	// 服务价格，精确到小数点后两位：单位:元
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 子订单实付金额。精确到2位小数，单位:元。如:200.07，表示:200元7分。
	Payment string `json:"payment,omitempty" xml:"payment,omitempty"`
	// 商品名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 服务子订单总费用
	TotalFee string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 最近退款的id
	RefundId string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 服务图片地址
	PicPath string `json:"pic_path,omitempty" xml:"pic_path,omitempty"`
	// 支持家装类物流的类型
	TmserSpuCode string `json:"tmser_spu_code,omitempty" xml:"tmser_spu_code,omitempty"`
	// 天猫汽车预约服务时间
	EtSerTime string `json:"et_ser_time,omitempty" xml:"et_ser_time,omitempty"`
	// 电子凭证预约门店地址
	EtShopName string `json:"et_shop_name,omitempty" xml:"et_shop_name,omitempty"`
	// 电子凭证核销门店地址
	EtVerifiedShopName string `json:"et_verified_shop_name,omitempty" xml:"et_verified_shop_name,omitempty"`
	// 车牌号
	EtPlateNumber string `json:"et_plate_number,omitempty" xml:"et_plate_number,omitempty"`
	// 虚拟服务子订单订单号(String类型)
	OidStr string `json:"oid_str,omitempty" xml:"oid_str,omitempty"`
	// appleCareEmail
	AppleCareEmail string `json:"apple_care_email,omitempty" xml:"apple_care_email,omitempty"`
	// appleCareMPN
	AppleCareMpn string `json:"apple_care_mpn,omitempty" xml:"apple_care_mpn,omitempty"`
	// 服务供应链-服务商外部编码标
	ServiceOuterId string `json:"service_outer_id,omitempty" xml:"service_outer_id,omitempty"`
	// 服务供应链-服务订单类型,1:主子挂载；2：双主挂载；3：单独售卖
	ServiceOrderType string `json:"service_order_type,omitempty" xml:"service_order_type,omitempty"`
	// 服务订单与实物订单关联关系
	ExtServiceBizId string `json:"ext_service_biz_id,omitempty" xml:"ext_service_biz_id,omitempty"`
	// 套餐购服务原价� (单位分)
	CommAmountUnit string `json:"comm_amount_unit,omitempty" xml:"comm_amount_unit,omitempty"`
	// 虚拟服务子订单订单号
	Oid int64 `json:"oid,omitempty" xml:"oid,omitempty"`
	// 服务所属的交易订单号。如果服务为一年包换，则item_oid这笔订单享受改服务的保护
	ItemOid int64 `json:"item_oid,omitempty" xml:"item_oid,omitempty"`
	// 服务数字id
	ServiceId int64 `json:"service_id,omitempty" xml:"service_id,omitempty"`
	// 购买数量，取值范围为大于0的整数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 套餐购对应的成分品id
	CombineSubItemId int64 `json:"combine_sub_item_id,omitempty" xml:"combine_sub_item_id,omitempty"`
	// 套餐购对应的成分品的skuId
	CombineSubItemSkuId int64 `json:"combine_sub_item_sku_id,omitempty" xml:"combine_sub_item_sku_id,omitempty"`
}

var poolServiceOrder = sync.Pool{
	New: func() any {
		return new(ServiceOrder)
	},
}

// GetServiceOrder() 从对象池中获取ServiceOrder
func GetServiceOrder() *ServiceOrder {
	return poolServiceOrder.Get().(*ServiceOrder)
}

// ReleaseServiceOrder 释放ServiceOrder
func ReleaseServiceOrder(v *ServiceOrder) {
	v.ServiceDetailUrl = ""
	v.Price = ""
	v.Payment = ""
	v.Title = ""
	v.TotalFee = ""
	v.BuyerNick = ""
	v.RefundId = ""
	v.SellerNick = ""
	v.PicPath = ""
	v.TmserSpuCode = ""
	v.EtSerTime = ""
	v.EtShopName = ""
	v.EtVerifiedShopName = ""
	v.EtPlateNumber = ""
	v.OidStr = ""
	v.AppleCareEmail = ""
	v.AppleCareMpn = ""
	v.ServiceOuterId = ""
	v.ServiceOrderType = ""
	v.ExtServiceBizId = ""
	v.CommAmountUnit = ""
	v.Oid = 0
	v.ItemOid = 0
	v.ServiceId = 0
	v.Num = 0
	v.CombineSubItemId = 0
	v.CombineSubItemSkuId = 0
	poolServiceOrder.Put(v)
}
