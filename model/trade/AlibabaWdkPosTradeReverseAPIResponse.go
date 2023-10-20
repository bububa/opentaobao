package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPosTradeReverseAPIResponse 轻pos品牌营销退款接口 API返回值
// alibaba.wdk.pos.trade.reverse
//
// 轻pos品牌营销场景，商家调用退款接口
type AlibabaWdkPosTradeReverseAPIResponse struct {
	model.CommonResponse
	AlibabaWdkPosTradeReverseAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkPosTradeReverseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkPosTradeReverseAPIResponseModel).Reset()
}

// AlibabaWdkPosTradeReverseAPIResponseModel is 轻pos品牌营销退款接口 成功返回结果
type AlibabaWdkPosTradeReverseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_pos_trade_reverse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款结果
	Result *FastBuyPosReverseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkPosTradeReverseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkPosTradeReverseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkPosTradeReverseAPIResponse)
	},
}

// GetAlibabaWdkPosTradeReverseAPIResponse 从 sync.Pool 获取 AlibabaWdkPosTradeReverseAPIResponse
func GetAlibabaWdkPosTradeReverseAPIResponse() *AlibabaWdkPosTradeReverseAPIResponse {
	return poolAlibabaWdkPosTradeReverseAPIResponse.Get().(*AlibabaWdkPosTradeReverseAPIResponse)
}

// ReleaseAlibabaWdkPosTradeReverseAPIResponse 将 AlibabaWdkPosTradeReverseAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkPosTradeReverseAPIResponse(v *AlibabaWdkPosTradeReverseAPIResponse) {
	v.Reset()
	poolAlibabaWdkPosTradeReverseAPIResponse.Put(v)
}
