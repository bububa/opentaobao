package tmallchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoChannelTradePrepayOfflineAddAPIResponse 渠道分销供应商上传线下流水预存款（增加） API返回值
// taobao.channel.trade.prepay.offline.add
//
// 渠道分销供应商上传线下流水预存款（增加）
type TaobaoChannelTradePrepayOfflineAddAPIResponse struct {
	model.CommonResponse
	TaobaoChannelTradePrepayOfflineAddAPIResponseModel
}

// TaobaoChannelTradePrepayOfflineAddAPIResponseModel is 渠道分销供应商上传线下流水预存款（增加） 成功返回结果
type TaobaoChannelTradePrepayOfflineAddAPIResponseModel struct {
	XMLName xml.Name `xml:"channel_trade_prepay_offline_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoChannelTradePrepayOfflineAddResultTopDo `json:"result,omitempty" xml:"result,omitempty"`
}
