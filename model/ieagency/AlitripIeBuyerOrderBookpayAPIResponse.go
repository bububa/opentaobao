package ieagency

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripiebuyerorderbookpayAPIResponse 【国际机票】下单预定支付 API返回值
// alitrip.ie.buyer.order.bookpay
//
// 【国际机票】 生单预定支付接口
type AlitripiebuyerorderbookpayAPIResponse struct {
	model.CommonResponse
	AlitripiebuyerorderbookpayAPIResponseModel
}

// AlitripiebuyerorderbookpayAPIResponseModel is 【国际机票】下单预定支付 成功返回结果
type AlitripiebuyerorderbookpayAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_buyer_order_bookpay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应
	BookPayOrderResult *BaseApiResult `json:"book_pay_order_result,omitempty" xml:"book_pay_order_result,omitempty"`
}
