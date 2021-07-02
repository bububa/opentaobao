package tmallchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallChannelTradeDeliverorderGetAPIResponse 通过发货单单号获取发货单的详情 API返回值
// tmall.channel.trade.deliverorder.get
//
// 通过发货单单号获取发货单的详情
type TmallChannelTradeDeliverorderGetAPIResponse struct {
	model.CommonResponse
	TmallChannelTradeDeliverorderGetAPIResponseModel
}

// TmallChannelTradeDeliverorderGetAPIResponseModel is 通过发货单单号获取发货单的详情 成功返回结果
type TmallChannelTradeDeliverorderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_channel_trade_deliverorder_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 发货单单号
	MainDeliverOrderNo int64 `json:"main_deliver_order_no,omitempty" xml:"main_deliver_order_no,omitempty"`
	// 发货单状态
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 确认单审核状态
	AuditStatus int64 `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 分销商Nick
	DistributorNick string `json:"distributor_nick,omitempty" xml:"distributor_nick,omitempty"`
	// 渠道
	Channel int64 `json:"channel,omitempty" xml:"channel,omitempty"`
	// 创建时间
	OrderCreateTime string `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
	// 最后更新时间
	OrderLastModifyTime string `json:"order_last_modify_time,omitempty" xml:"order_last_modify_time,omitempty"`
	// 子发货单列表
	SubDeliverOrderList []TopChannelSubDeliverOrderDto `json:"sub_deliver_order_list,omitempty" xml:"sub_deliver_order_list>top_channel_sub_deliver_order_dto,omitempty"`
	// 物流单列表
	LogisticsOrderList []TopChannelLogisticsOrderDto `json:"logistics_order_list,omitempty" xml:"logistics_order_list>top_channel_logistics_order_dto,omitempty"`
	// 解析描述
	Schema string `json:"schema,omitempty" xml:"schema,omitempty"`
}
