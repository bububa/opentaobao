package refund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundNegotiatereturnAPIResponse 协商退货退款 API返回值
// taobao.refund.negotiatereturn
//
// 协商退货退款
type TaobaoRefundNegotiatereturnAPIResponse struct {
	model.CommonResponse
	TaobaoRefundNegotiatereturnAPIResponseModel
}

// TaobaoRefundNegotiatereturnAPIResponseModel is 协商退货退款 成功返回结果
type TaobaoRefundNegotiatereturnAPIResponseModel struct {
	XMLName xml.Name `xml:"refund_negotiatereturn_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
