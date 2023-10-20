package ieagency

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripIeBuyerOrderBookpayAPIResponse 【国际机票】下单预定支付 API返回值
// alitrip.ie.buyer.order.bookpay
//
// 【国际机票】 生单预定支付接口
type AlitripIeBuyerOrderBookpayAPIResponse struct {
	model.CommonResponse
	AlitripIeBuyerOrderBookpayAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripIeBuyerOrderBookpayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripIeBuyerOrderBookpayAPIResponseModel).Reset()
}

// AlitripIeBuyerOrderBookpayAPIResponseModel is 【国际机票】下单预定支付 成功返回结果
type AlitripIeBuyerOrderBookpayAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_buyer_order_bookpay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应
	BookPayOrderResult *BaseApiResult `json:"book_pay_order_result,omitempty" xml:"book_pay_order_result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripIeBuyerOrderBookpayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BookPayOrderResult = nil
}

var poolAlitripIeBuyerOrderBookpayAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripIeBuyerOrderBookpayAPIResponse)
	},
}

// GetAlitripIeBuyerOrderBookpayAPIResponse 从 sync.Pool 获取 AlitripIeBuyerOrderBookpayAPIResponse
func GetAlitripIeBuyerOrderBookpayAPIResponse() *AlitripIeBuyerOrderBookpayAPIResponse {
	return poolAlitripIeBuyerOrderBookpayAPIResponse.Get().(*AlitripIeBuyerOrderBookpayAPIResponse)
}

// ReleaseAlitripIeBuyerOrderBookpayAPIResponse 将 AlitripIeBuyerOrderBookpayAPIResponse 保存到 sync.Pool
func ReleaseAlitripIeBuyerOrderBookpayAPIResponse(v *AlitripIeBuyerOrderBookpayAPIResponse) {
	v.Reset()
	poolAlitripIeBuyerOrderBookpayAPIResponse.Put(v)
}
