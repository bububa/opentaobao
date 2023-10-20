package tbrefund

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoRefundNegotiatereturnAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRefundNegotiatereturnAPIResponseModel).Reset()
}

// TaobaoRefundNegotiatereturnAPIResponseModel is 协商退货退款 成功返回结果
type TaobaoRefundNegotiatereturnAPIResponseModel struct {
	XMLName xml.Name `xml:"refund_negotiatereturn_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRefundNegotiatereturnAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolTaobaoRefundNegotiatereturnAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRefundNegotiatereturnAPIResponse)
	},
}

// GetTaobaoRefundNegotiatereturnAPIResponse 从 sync.Pool 获取 TaobaoRefundNegotiatereturnAPIResponse
func GetTaobaoRefundNegotiatereturnAPIResponse() *TaobaoRefundNegotiatereturnAPIResponse {
	return poolTaobaoRefundNegotiatereturnAPIResponse.Get().(*TaobaoRefundNegotiatereturnAPIResponse)
}

// ReleaseTaobaoRefundNegotiatereturnAPIResponse 将 TaobaoRefundNegotiatereturnAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRefundNegotiatereturnAPIResponse(v *TaobaoRefundNegotiatereturnAPIResponse) {
	v.Reset()
	poolTaobaoRefundNegotiatereturnAPIResponse.Put(v)
}
