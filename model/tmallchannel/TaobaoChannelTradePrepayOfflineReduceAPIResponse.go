package tmallchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoChannelTradePrepayOfflineReduceAPIResponse
渠道分销供应商上传线下流水预存款（减少） API返回值
taobao.channel.trade.prepay.offline.reduce

渠道分销供应商上传线下流水预存款（减少） */
type TaobaoChannelTradePrepayOfflineReduceAPIResponse struct {
	model.CommonResponse
	TaobaoChannelTradePrepayOfflineReduceAPIResponseModel
}

// TaobaoChannelTradePrepayOfflineReduceAPIResponseModel is 渠道分销供应商上传线下流水预存款（减少） 成功返回结果
type TaobaoChannelTradePrepayOfflineReduceAPIResponseModel struct {
	XMLName xml.Name `xml:"channel_trade_prepay_offline_reduce_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoChannelTradePrepayOfflineReduceResultTopDo `json:"result,omitempty" xml:"result,omitempty"`
}
