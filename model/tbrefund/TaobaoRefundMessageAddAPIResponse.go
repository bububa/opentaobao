package tbrefund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundMessageAddAPIResponse 创建退款留言/凭证 API返回值
// taobao.refund.message.add
//
// 创建退款留言/凭证
type TaobaoRefundMessageAddAPIResponse struct {
	model.CommonResponse
	TaobaoRefundMessageAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRefundMessageAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRefundMessageAddAPIResponseModel).Reset()
}

// TaobaoRefundMessageAddAPIResponseModel is 创建退款留言/凭证 成功返回结果
type TaobaoRefundMessageAddAPIResponseModel struct {
	XMLName xml.Name `xml:"refund_message_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款信息。包含id和created
	RefundMessage *RefundMessage `json:"refund_message,omitempty" xml:"refund_message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRefundMessageAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RefundMessage = nil
}

var poolTaobaoRefundMessageAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRefundMessageAddAPIResponse)
	},
}

// GetTaobaoRefundMessageAddAPIResponse 从 sync.Pool 获取 TaobaoRefundMessageAddAPIResponse
func GetTaobaoRefundMessageAddAPIResponse() *TaobaoRefundMessageAddAPIResponse {
	return poolTaobaoRefundMessageAddAPIResponse.Get().(*TaobaoRefundMessageAddAPIResponse)
}

// ReleaseTaobaoRefundMessageAddAPIResponse 将 TaobaoRefundMessageAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRefundMessageAddAPIResponse(v *TaobaoRefundMessageAddAPIResponse) {
	v.Reset()
	poolTaobaoRefundMessageAddAPIResponse.Put(v)
}
