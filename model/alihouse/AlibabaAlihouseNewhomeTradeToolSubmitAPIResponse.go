package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeTradeToolSubmitAPIResponse 交易工具信息上翻 API返回值
// alibaba.alihouse.newhome.trade.tool.submit
//
// 交易工具信息上翻
type AlibabaAlihouseNewhomeTradeToolSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeTradeToolSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeTradeToolSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeTradeToolSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeTradeToolSubmitAPIResponseModel is 交易工具信息上翻 成功返回结果
type AlibabaAlihouseNewhomeTradeToolSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_trade_tool_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaAlihouseNewhomeTradeToolSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeTradeToolSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeTradeToolSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeTradeToolSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeTradeToolSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeTradeToolSubmitAPIResponse
func GetAlibabaAlihouseNewhomeTradeToolSubmitAPIResponse() *AlibabaAlihouseNewhomeTradeToolSubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeTradeToolSubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeTradeToolSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeTradeToolSubmitAPIResponse 将 AlibabaAlihouseNewhomeTradeToolSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeTradeToolSubmitAPIResponse(v *AlibabaAlihouseNewhomeTradeToolSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeTradeToolSubmitAPIResponse.Put(v)
}
