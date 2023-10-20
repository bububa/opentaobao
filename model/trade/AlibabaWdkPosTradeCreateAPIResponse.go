package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPosTradeCreateAPIResponse 轻pos品牌营销下单接口 API返回值
// alibaba.wdk.pos.trade.create
//
// 提供给石基进行轻pos品牌营销下单
type AlibabaWdkPosTradeCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkPosTradeCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkPosTradeCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkPosTradeCreateAPIResponseModel).Reset()
}

// AlibabaWdkPosTradeCreateAPIResponseModel is 轻pos品牌营销下单接口 成功返回结果
type AlibabaWdkPosTradeCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_pos_trade_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创单结果
	Result *FastBuyPosCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkPosTradeCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkPosTradeCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkPosTradeCreateAPIResponse)
	},
}

// GetAlibabaWdkPosTradeCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkPosTradeCreateAPIResponse
func GetAlibabaWdkPosTradeCreateAPIResponse() *AlibabaWdkPosTradeCreateAPIResponse {
	return poolAlibabaWdkPosTradeCreateAPIResponse.Get().(*AlibabaWdkPosTradeCreateAPIResponse)
}

// ReleaseAlibabaWdkPosTradeCreateAPIResponse 将 AlibabaWdkPosTradeCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkPosTradeCreateAPIResponse(v *AlibabaWdkPosTradeCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkPosTradeCreateAPIResponse.Put(v)
}
