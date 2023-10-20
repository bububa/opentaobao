package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeTradeToolBindAPIResponse 批量绑定交易工具 API返回值
// alibaba.alihouse.newhome.trade.tool.bind
//
// 批量绑定交易工具
type AlibabaAlihouseNewhomeTradeToolBindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeTradeToolBindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeTradeToolBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeTradeToolBindAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeTradeToolBindAPIResponseModel is 批量绑定交易工具 成功返回结果
type AlibabaAlihouseNewhomeTradeToolBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_trade_tool_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 对象
	Result *AlibabaAlihouseNewhomeTradeToolBindResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeTradeToolBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeTradeToolBindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeTradeToolBindAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeTradeToolBindAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeTradeToolBindAPIResponse
func GetAlibabaAlihouseNewhomeTradeToolBindAPIResponse() *AlibabaAlihouseNewhomeTradeToolBindAPIResponse {
	return poolAlibabaAlihouseNewhomeTradeToolBindAPIResponse.Get().(*AlibabaAlihouseNewhomeTradeToolBindAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeTradeToolBindAPIResponse 将 AlibabaAlihouseNewhomeTradeToolBindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeTradeToolBindAPIResponse(v *AlibabaAlihouseNewhomeTradeToolBindAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeTradeToolBindAPIResponse.Put(v)
}
