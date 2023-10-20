package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPosTradeCloseAPIResponse 轻pos品牌营销关单接口 API返回值
// alibaba.wdk.pos.trade.close
//
// 轻pos品牌营销场景，提供关单接口给外部商家
type AlibabaWdkPosTradeCloseAPIResponse struct {
	model.CommonResponse
	AlibabaWdkPosTradeCloseAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkPosTradeCloseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkPosTradeCloseAPIResponseModel).Reset()
}

// AlibabaWdkPosTradeCloseAPIResponseModel is 轻pos品牌营销关单接口 成功返回结果
type AlibabaWdkPosTradeCloseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_pos_trade_close_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关单结果
	Result *FastBuyPosCloseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkPosTradeCloseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkPosTradeCloseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkPosTradeCloseAPIResponse)
	},
}

// GetAlibabaWdkPosTradeCloseAPIResponse 从 sync.Pool 获取 AlibabaWdkPosTradeCloseAPIResponse
func GetAlibabaWdkPosTradeCloseAPIResponse() *AlibabaWdkPosTradeCloseAPIResponse {
	return poolAlibabaWdkPosTradeCloseAPIResponse.Get().(*AlibabaWdkPosTradeCloseAPIResponse)
}

// ReleaseAlibabaWdkPosTradeCloseAPIResponse 将 AlibabaWdkPosTradeCloseAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkPosTradeCloseAPIResponse(v *AlibabaWdkPosTradeCloseAPIResponse) {
	v.Reset()
	poolAlibabaWdkPosTradeCloseAPIResponse.Put(v)
}
