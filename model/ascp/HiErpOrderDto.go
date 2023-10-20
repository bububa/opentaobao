package ascp

import (
	"sync"
)

// HiErpOrderDto 结构体
type HiErpOrderDto struct {
	// 外部订单号，唯一标识
	OutOrderCode string `json:"out_order_code,omitempty" xml:"out_order_code,omitempty"`
	// 淘系交易ID
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 渠道，比如213
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 店铺名称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 外部店铺名称
	OutSellerNick string `json:"out_seller_nick,omitempty" xml:"out_seller_nick,omitempty"`
	// 订单类型，目前只支持销售单，取值为0
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 下单时间，为空则取接口调用时间
	OrderTime string `json:"order_time,omitempty" xml:"order_time,omitempty"`
	// 支付时间，为空则取接口调用时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 买家留言
	BuyerMemo string `json:"buyer_memo,omitempty" xml:"buyer_memo,omitempty"`
	// 扩展字段，json格式
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 卖家备注
	SellerMessage string `json:"seller_message,omitempty" xml:"seller_message,omitempty"`
	// 仓编码， 揽配业务必填
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 货主ID
	OwnerId int64 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
	// 是否加密，默认为false
	CipherText bool `json:"cipher_text,omitempty" xml:"cipher_text,omitempty"`
	// 是否自动下发，默认为false
	AutoSend bool `json:"auto_send,omitempty" xml:"auto_send,omitempty"`
}

var poolHiErpOrderDto = sync.Pool{
	New: func() any {
		return new(HiErpOrderDto)
	},
}

// GetHiErpOrderDto() 从对象池中获取HiErpOrderDto
func GetHiErpOrderDto() *HiErpOrderDto {
	return poolHiErpOrderDto.Get().(*HiErpOrderDto)
}

// ReleaseHiErpOrderDto 释放HiErpOrderDto
func ReleaseHiErpOrderDto(v *HiErpOrderDto) {
	v.OutOrderCode = ""
	v.TradeId = ""
	v.Channel = ""
	v.SellerNick = ""
	v.OutSellerNick = ""
	v.OrderType = ""
	v.OrderTime = ""
	v.PayTime = ""
	v.BuyerMemo = ""
	v.Feature = ""
	v.SellerMessage = ""
	v.StoreCode = ""
	v.OwnerId = 0
	v.CipherText = false
	v.AutoSend = false
	poolHiErpOrderDto.Put(v)
}
