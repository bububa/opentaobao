package tmallchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallChannelTradeApplyorderAgreeAPIResponse 供应商审核同意采购申请单 API返回值
// tmall.channel.trade.applyorder.agree
//
// 供应商审核同意采购申请单
type TmallChannelTradeApplyorderAgreeAPIResponse struct {
	model.CommonResponse
	TmallChannelTradeApplyorderAgreeAPIResponseModel
}

// TmallChannelTradeApplyorderAgreeAPIResponseModel is 供应商审核同意采购申请单 成功返回结果
type TmallChannelTradeApplyorderAgreeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_channel_trade_applyorder_agree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallChannelTradeApplyorderAgreeResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
