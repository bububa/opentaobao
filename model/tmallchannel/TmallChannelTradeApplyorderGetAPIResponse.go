package tmallchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallChannelTradeApplyorderGetAPIResponse 查询采购申请单详情 API返回值
// tmall.channel.trade.applyorder.get
//
// 通过采购申请单ID获取单据详情
type TmallChannelTradeApplyorderGetAPIResponse struct {
	model.CommonResponse
	TmallChannelTradeApplyorderGetAPIResponseModel
}

// TmallChannelTradeApplyorderGetAPIResponseModel is 查询采购申请单详情 成功返回结果
type TmallChannelTradeApplyorderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_channel_trade_applyorder_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 申请单单号
	ChannelPurchaseApplyOrderNo string `json:"channel_purchase_apply_order_no,omitempty" xml:"channel_purchase_apply_order_no,omitempty"`
	// 分销商名称
	DistributorNick string `json:"distributor_nick,omitempty" xml:"distributor_nick,omitempty"`
	// 审核描述
	AuditDesc string `json:"audit_desc,omitempty" xml:"audit_desc,omitempty"`
	// 撤回的描述
	CancelDesc string `json:"cancel_desc,omitempty" xml:"cancel_desc,omitempty"`
	// 审核创建时间
	ApplyCreateTime string `json:"apply_create_time,omitempty" xml:"apply_create_time,omitempty"`
	// 解析详情
	Schema string `json:"schema,omitempty" xml:"schema,omitempty"`
	// 渠道 : com.tmall.usc.support.common.channelcenter.channel.enums.Channel
	Channel int64 `json:"channel,omitempty" xml:"channel,omitempty"`
	// 审核状态 : com.tmall.usc.support.common.channelcenter.trade.enums.AuditStatus
	AuditStatus int64 `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 申请单状态 : com.tmall.usc.support.common.channelcenter.trade.enums.apply.ChannelPurchaseApplyOrderStatus
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 交易类型: com.tmall.usc.support.common.channelcenter.trade.enums.TradeType
	TradeType int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
	// 申请单详情
	ApplyOrderDetail *TopChannelApplyOrderDetailDto `json:"apply_order_detail,omitempty" xml:"apply_order_detail,omitempty"`
}
