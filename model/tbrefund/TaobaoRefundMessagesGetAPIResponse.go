package tbrefund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaorefundmessagesgetAPIResponse 查询退款留言/凭证列表 API返回值
// taobao.refund.messages.get
//
// 查询退款留言/凭证列表
type TaobaorefundmessagesgetAPIResponse struct {
	model.CommonResponse
	TaobaorefundmessagesgetAPIResponseModel
}

// TaobaorefundmessagesgetAPIResponseModel is 查询退款留言/凭证列表 成功返回结果
type TaobaorefundmessagesgetAPIResponseModel struct {
	XMLName xml.Name `xml:"refund_messages_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询到的退款留言/凭证列表
	RefundMessages []RefundMessage `json:"refund_messages,omitempty" xml:"refund_messages>refund_message,omitempty"`
	// 查询到的退款留言/凭证总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
