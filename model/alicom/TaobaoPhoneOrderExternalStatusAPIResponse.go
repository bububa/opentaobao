package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPhoneOrderExternalStatusAPIResponse 话费外放订单状态接口 API返回值
// taobao.phone.order.external.status
//
// 话费外放订单状态接口
type TaobaoPhoneOrderExternalStatusAPIResponse struct {
	model.CommonResponse
	TaobaoPhoneOrderExternalStatusAPIResponseModel
}

// TaobaoPhoneOrderExternalStatusAPIResponseModel is 话费外放订单状态接口 成功返回结果
type TaobaoPhoneOrderExternalStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"phone_order_external_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应编码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 响应描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 返回结果
	Model *DistributeTradeOrderInfo `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
