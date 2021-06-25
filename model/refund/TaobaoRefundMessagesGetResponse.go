package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询退款留言/凭证列表 APIResponse
taobao.refund.messages.get

查询退款留言/凭证列表
*/
type TaobaoRefundMessagesGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoRefundMessagesGetResponse `json:"taobao_refund_messages_get_response,omitempty"`
}

type TaobaoRefundMessagesGetResponse struct {

    // 查询到的退款留言/凭证总数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 查询到的退款留言/凭证列表
    RefundMessages   []RefundMessage `json:"refund_messages,omitempty"`

}
