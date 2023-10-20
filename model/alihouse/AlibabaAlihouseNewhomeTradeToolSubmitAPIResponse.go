package alihouse

import (
	"encoding/xml"

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

// AlibabaAlihouseNewhomeTradeToolSubmitAPIResponseModel is 交易工具信息上翻 成功返回结果
type AlibabaAlihouseNewhomeTradeToolSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_trade_tool_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaAlihouseNewhomeTradeToolSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
