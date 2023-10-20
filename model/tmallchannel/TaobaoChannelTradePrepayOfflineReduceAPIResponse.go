package tmallchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoChannelTradePrepayOfflineReduceAPIResponse 渠道分销供应商上传线下流水预存款（减少） API返回值
// taobao.channel.trade.prepay.offline.reduce
//
// 渠道分销供应商上传线下流水预存款（减少）
type TaobaoChannelTradePrepayOfflineReduceAPIResponse struct {
	model.CommonResponse
	TaobaoChannelTradePrepayOfflineReduceAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoChannelTradePrepayOfflineReduceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoChannelTradePrepayOfflineReduceAPIResponseModel).Reset()
}

// TaobaoChannelTradePrepayOfflineReduceAPIResponseModel is 渠道分销供应商上传线下流水预存款（减少） 成功返回结果
type TaobaoChannelTradePrepayOfflineReduceAPIResponseModel struct {
	XMLName xml.Name `xml:"channel_trade_prepay_offline_reduce_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoChannelTradePrepayOfflineReduceResultTopDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoChannelTradePrepayOfflineReduceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoChannelTradePrepayOfflineReduceAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoChannelTradePrepayOfflineReduceAPIResponse)
	},
}

// GetTaobaoChannelTradePrepayOfflineReduceAPIResponse 从 sync.Pool 获取 TaobaoChannelTradePrepayOfflineReduceAPIResponse
func GetTaobaoChannelTradePrepayOfflineReduceAPIResponse() *TaobaoChannelTradePrepayOfflineReduceAPIResponse {
	return poolTaobaoChannelTradePrepayOfflineReduceAPIResponse.Get().(*TaobaoChannelTradePrepayOfflineReduceAPIResponse)
}

// ReleaseTaobaoChannelTradePrepayOfflineReduceAPIResponse 将 TaobaoChannelTradePrepayOfflineReduceAPIResponse 保存到 sync.Pool
func ReleaseTaobaoChannelTradePrepayOfflineReduceAPIResponse(v *TaobaoChannelTradePrepayOfflineReduceAPIResponse) {
	v.Reset()
	poolTaobaoChannelTradePrepayOfflineReduceAPIResponse.Put(v)
}
