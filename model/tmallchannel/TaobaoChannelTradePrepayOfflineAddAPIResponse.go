package tmallchannel

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoChannelTradePrepayOfflineAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoChannelTradePrepayOfflineAddAPIResponseModel).Reset()
}

// TaobaoChannelTradePrepayOfflineAddAPIResponseModel is 渠道分销供应商上传线下流水预存款（增加） 成功返回结果
type TaobaoChannelTradePrepayOfflineAddAPIResponseModel struct {
	XMLName xml.Name `xml:"channel_trade_prepay_offline_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoChannelTradePrepayOfflineAddResultTopDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoChannelTradePrepayOfflineAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoChannelTradePrepayOfflineAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoChannelTradePrepayOfflineAddAPIResponse)
	},
}

// GetTaobaoChannelTradePrepayOfflineAddAPIResponse 从 sync.Pool 获取 TaobaoChannelTradePrepayOfflineAddAPIResponse
func GetTaobaoChannelTradePrepayOfflineAddAPIResponse() *TaobaoChannelTradePrepayOfflineAddAPIResponse {
	return poolTaobaoChannelTradePrepayOfflineAddAPIResponse.Get().(*TaobaoChannelTradePrepayOfflineAddAPIResponse)
}

// ReleaseTaobaoChannelTradePrepayOfflineAddAPIResponse 将 TaobaoChannelTradePrepayOfflineAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoChannelTradePrepayOfflineAddAPIResponse(v *TaobaoChannelTradePrepayOfflineAddAPIResponse) {
	v.Reset()
	poolTaobaoChannelTradePrepayOfflineAddAPIResponse.Put(v)
}
