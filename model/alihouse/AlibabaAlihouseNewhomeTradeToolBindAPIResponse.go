package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhometradetoolbindAPIResponse 批量绑定交易工具 API返回值
// alibaba.alihouse.newhome.trade.tool.bind
//
// 批量绑定交易工具
type AlibabaalihousenewhometradetoolbindAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhometradetoolbindAPIResponseModel
}

// AlibabaalihousenewhometradetoolbindAPIResponseModel is 批量绑定交易工具 成功返回结果
type AlibabaalihousenewhometradetoolbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_trade_tool_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 对象
	Result *AlibabaalihousenewhometradetoolbindResult `json:"result,omitempty" xml:"result,omitempty"`
}
