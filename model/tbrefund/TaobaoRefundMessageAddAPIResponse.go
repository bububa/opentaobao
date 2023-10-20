package tbrefund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaorefundmessageaddAPIResponse 创建退款留言/凭证 API返回值
// taobao.refund.message.add
//
// 创建退款留言/凭证
type TaobaorefundmessageaddAPIResponse struct {
	model.CommonResponse
	TaobaorefundmessageaddAPIResponseModel
}

// TaobaorefundmessageaddAPIResponseModel is 创建退款留言/凭证 成功返回结果
type TaobaorefundmessageaddAPIResponseModel struct {
	XMLName xml.Name `xml:"refund_message_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款信息。包含id和created
	RefundMessage *RefundMessage `json:"refund_message,omitempty" xml:"refund_message,omitempty"`
}
