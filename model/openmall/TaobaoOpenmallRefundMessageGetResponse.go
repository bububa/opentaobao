package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
openmall获取退款单留言 API返回值 
taobao.openmall.refund.message.get

openmall获取退款单留言
*/
type TaobaoOpenmallRefundMessageGetAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallRefundMessageGetResponse
}

// openmall获取退款单留言 成功返回结果
type TaobaoOpenmallRefundMessageGetResponse struct {
    XMLName xml.Name `xml:"openmall_refund_message_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 留言列表
    ResultsList   []RefundMessage `json:"results_list,omitempty" xml:"results_list>refund_message,omitempty"`
}
