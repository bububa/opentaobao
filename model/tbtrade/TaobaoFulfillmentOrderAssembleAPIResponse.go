package tbtrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFulfillmentOrderAssembleAPIResponse 拆合单结果回传接口 API返回值
// taobao.fulfillment.order.assemble
//
// 拆合单结果回传接口
type TaobaoFulfillmentOrderAssembleAPIResponse struct {
	model.CommonResponse
	TaobaoFulfillmentOrderAssembleAPIResponseModel
}

// TaobaoFulfillmentOrderAssembleAPIResponseModel is 拆合单结果回传接口 成功返回结果
type TaobaoFulfillmentOrderAssembleAPIResponseModel struct {
	XMLName xml.Name `xml:"fulfillment_order_assemble_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	CallErrorCode string `json:"call_error_code,omitempty" xml:"call_error_code,omitempty"`
	// 错误信息描述
	CallErrorMsg string `json:"call_error_msg,omitempty" xml:"call_error_msg,omitempty"`
	// 回传结果
	Model *OrderAssembleResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 调用结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
