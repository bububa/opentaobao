package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPosTradePayAPIResponse 轻pos品牌营销支付接口 API返回值
// alibaba.wdk.pos.trade.pay
//
// 轻pos场景，外部商家支付后调用开放平台把支付信息回传给五道口交易
type AlibabaWdkPosTradePayAPIResponse struct {
	model.CommonResponse
	AlibabaWdkPosTradePayAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkPosTradePayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkPosTradePayAPIResponseModel).Reset()
}

// AlibabaWdkPosTradePayAPIResponseModel is 轻pos品牌营销支付接口 成功返回结果
type AlibabaWdkPosTradePayAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_pos_trade_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 支付结果
	Result *FastBuyPosPayResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkPosTradePayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkPosTradePayAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkPosTradePayAPIResponse)
	},
}

// GetAlibabaWdkPosTradePayAPIResponse 从 sync.Pool 获取 AlibabaWdkPosTradePayAPIResponse
func GetAlibabaWdkPosTradePayAPIResponse() *AlibabaWdkPosTradePayAPIResponse {
	return poolAlibabaWdkPosTradePayAPIResponse.Get().(*AlibabaWdkPosTradePayAPIResponse)
}

// ReleaseAlibabaWdkPosTradePayAPIResponse 将 AlibabaWdkPosTradePayAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkPosTradePayAPIResponse(v *AlibabaWdkPosTradePayAPIResponse) {
	v.Reset()
	poolAlibabaWdkPosTradePayAPIResponse.Put(v)
}
