package tmallchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallChannelTradeDeliverorderRejectAPIResponse 供应商拒绝收货确认单 API返回值
// tmall.channel.trade.deliverorder.reject
//
// 供应商拒绝收货确认单
type TmallChannelTradeDeliverorderRejectAPIResponse struct {
	model.CommonResponse
	TmallChannelTradeDeliverorderRejectAPIResponseModel
}

// TmallChannelTradeDeliverorderRejectAPIResponseModel is 供应商拒绝收货确认单 成功返回结果
type TmallChannelTradeDeliverorderRejectAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_channel_trade_deliverorder_reject_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallChannelTradeDeliverorderRejectResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
