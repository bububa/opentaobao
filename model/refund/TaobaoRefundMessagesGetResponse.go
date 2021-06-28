package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询退款留言/凭证列表 APIResponse
taobao.refund.messages.get

查询退款留言/凭证列表
*/
type TaobaoRefundMessagesGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"refund_messages_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询到的退款留言/凭证总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"