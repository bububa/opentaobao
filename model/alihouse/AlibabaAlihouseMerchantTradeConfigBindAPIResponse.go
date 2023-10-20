package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseMerchantTradeConfigBindAPIResponse 交易场景绑定 API返回值
// alibaba.alihouse.merchant.trade.config.bind
//
// 交易场景绑定
type AlibabaAlihouseMerchantTradeConfigBindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseMerchantTradeConfigBindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseMerchantTradeConfigBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseMerchantTradeConfigBindAPIResponseModel).Reset()
}

// AlibabaAlihouseMerchantTradeConfigBindAPIResponseModel is 交易场景绑定 成功返回结果
type AlibabaAlihouseMerchantTradeConfigBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_merchant_trade_config_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseMerchantTradeConfigBindResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseMerchantTradeConfigBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseMerchantTradeConfigBindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseMerchantTradeConfigBindAPIResponse)
	},
}

// GetAlibabaAlihouseMerchantTradeConfigBindAPIResponse 从 sync.Pool 获取 AlibabaAlihouseMerchantTradeConfigBindAPIResponse
func GetAlibabaAlihouseMerchantTradeConfigBindAPIResponse() *AlibabaAlihouseMerchantTradeConfigBindAPIResponse {
	return poolAlibabaAlihouseMerchantTradeConfigBindAPIResponse.Get().(*AlibabaAlihouseMerchantTradeConfigBindAPIResponse)
}

// ReleaseAlibabaAlihouseMerchantTradeConfigBindAPIResponse 将 AlibabaAlihouseMerchantTradeConfigBindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseMerchantTradeConfigBindAPIResponse(v *AlibabaAlihouseMerchantTradeConfigBindAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseMerchantTradeConfigBindAPIResponse.Put(v)
}
