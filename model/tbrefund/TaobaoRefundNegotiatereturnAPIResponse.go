package tbrefund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaorefundnegotiatereturnAPIResponse 协商退货退款 API返回值
// taobao.refund.negotiatereturn
//
// 协商退货退款
type TaobaorefundnegotiatereturnAPIResponse struct {
	model.CommonResponse
	TaobaorefundnegotiatereturnAPIResponseModel
}

// TaobaorefundnegotiatereturnAPIResponseModel is 协商退货退款 成功返回结果
type TaobaorefundnegotiatereturnAPIResponseModel struct {
	XMLName xml.Name `xml:"refund_negotiatereturn_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
