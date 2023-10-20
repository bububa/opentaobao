package tbrefund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundMessagesGetAPIResponse 查询退款留言/凭证列表 API返回值
// taobao.refund.messages.get
//
// 查询退款留言/凭证列表
type TaobaoRefundMessagesGetAPIResponse struct {
	model.CommonResponse
	TaobaoRefundMessagesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRefundMessagesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRefundMessagesGetAPIResponseModel).Reset()
}

// TaobaoRefundMessagesGetAPIResponseModel is 查询退款留言/凭证列表 成功返回结果
type TaobaoRefundMessagesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"refund_messages_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询到的退款留言/凭证列表
	RefundMessages []RefundMessage `json:"refund_messages,omitempty" xml:"refund_messages>refund_message,omitempty"`
	// 查询到的退款留言/凭证总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRefundMessagesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RefundMessages = m.RefundMessages[:0]
	m.TotalResults = 0
}

var poolTaobaoRefundMessagesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRefundMessagesGetAPIResponse)
	},
}

// GetTaobaoRefundMessagesGetAPIResponse 从 sync.Pool 获取 TaobaoRefundMessagesGetAPIResponse
func GetTaobaoRefundMessagesGetAPIResponse() *TaobaoRefundMessagesGetAPIResponse {
	return poolTaobaoRefundMessagesGetAPIResponse.Get().(*TaobaoRefundMessagesGetAPIResponse)
}

// ReleaseTaobaoRefundMessagesGetAPIResponse 将 TaobaoRefundMessagesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRefundMessagesGetAPIResponse(v *TaobaoRefundMessagesGetAPIResponse) {
	v.Reset()
	poolTaobaoRefundMessagesGetAPIResponse.Put(v)
}
