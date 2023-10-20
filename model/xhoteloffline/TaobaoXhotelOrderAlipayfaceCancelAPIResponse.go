package xhoteloffline

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelorderalipayfacecancelAPIResponse 线下信用住订单取消 API返回值
// taobao.xhotel.order.alipayface.cancel
//
// 提供给卖家进行线下信用住的订单取消。此接口仅仅支持线下信用住订单的取消
type TaobaoxhotelorderalipayfacecancelAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelorderalipayfacecancelAPIResponseModel
}

// TaobaoxhotelorderalipayfacecancelAPIResponseModel is 线下信用住订单取消 成功返回结果
type TaobaoxhotelorderalipayfacecancelAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_alipayface_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
