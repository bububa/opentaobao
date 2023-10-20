package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTradeAlianceCreateAPIResponse 推客平台订单回流 API返回值
// alibaba.trade.aliance.create
//
// 推客平台订单回流
type AlibabaTradeAlianceCreateAPIResponse struct {
	model.CommonResponse
	AlibabaTradeAlianceCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTradeAlianceCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTradeAlianceCreateAPIResponseModel).Reset()
}

// AlibabaTradeAlianceCreateAPIResponseModel is 推客平台订单回流 成功返回结果
type AlibabaTradeAlianceCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_trade_aliance_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单创建结果
	Result *AlibabaTradeAlianceCreateResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTradeAlianceCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTradeAlianceCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTradeAlianceCreateAPIResponse)
	},
}

// GetAlibabaTradeAlianceCreateAPIResponse 从 sync.Pool 获取 AlibabaTradeAlianceCreateAPIResponse
func GetAlibabaTradeAlianceCreateAPIResponse() *AlibabaTradeAlianceCreateAPIResponse {
	return poolAlibabaTradeAlianceCreateAPIResponse.Get().(*AlibabaTradeAlianceCreateAPIResponse)
}

// ReleaseAlibabaTradeAlianceCreateAPIResponse 将 AlibabaTradeAlianceCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTradeAlianceCreateAPIResponse(v *AlibabaTradeAlianceCreateAPIResponse) {
	v.Reset()
	poolAlibabaTradeAlianceCreateAPIResponse.Put(v)
}
