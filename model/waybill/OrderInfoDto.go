package waybill

import (
	"sync"
)

// OrderInfoDto 结构体
type OrderInfoDto struct {
	// 订单号,数量限制100，订单号（只限传入数字、字母、下划线和中划线，为避免出现冲突，请按电商平台真实订单号传入，请避免使用同个订单号重复取号）
	TradeOrderList []string `json:"trade_order_list,omitempty" xml:"trade_order_list>string,omitempty"`
	// 外部电商平台交易单号集合，非必填，数量限制100
	OutTradeOrderList []string `json:"out_trade_order_list,omitempty" xml:"out_trade_order_list>string,omitempty"`
	// 外部电商平台交易子单号集合，非必填，数量限制100
	OutTradeSubOrderList []string `json:"out_trade_sub_order_list,omitempty" xml:"out_trade_sub_order_list>string,omitempty"`
	// &lt;a href=&#34;http://open.taobao.com/doc2/detail.htm?spm=a219a.7629140.0.0.8cf9Nj&amp;treeId=17&amp;articleId=105085&amp;docType=1#2&#34;&gt;订单渠道平台编码&lt;/a&gt;
	OrderChannelsType string `json:"order_channels_type,omitempty" xml:"order_channels_type,omitempty"`
}

var poolOrderInfoDto = sync.Pool{
	New: func() any {
		return new(OrderInfoDto)
	},
}

// GetOrderInfoDto() 从对象池中获取OrderInfoDto
func GetOrderInfoDto() *OrderInfoDto {
	return poolOrderInfoDto.Get().(*OrderInfoDto)
}

// ReleaseOrderInfoDto 释放OrderInfoDto
func ReleaseOrderInfoDto(v *OrderInfoDto) {
	v.TradeOrderList = v.TradeOrderList[:0]
	v.OutTradeOrderList = v.OutTradeOrderList[:0]
	v.OutTradeSubOrderList = v.OutTradeSubOrderList[:0]
	v.OrderChannelsType = ""
	poolOrderInfoDto.Put(v)
}
