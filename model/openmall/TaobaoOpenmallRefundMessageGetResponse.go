package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
openmall获取退款单留言 APIResponse
taobao.openmall.refund.message.get

openmall获取退款单留言
*/
type TaobaoOpenmallRefundMessageGetAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallRefundMessageGetResponse
}

type TaobaoOpenmallRefundMessageGetResponse struct {
    XMLName xml.Name `xml:"openmall_refund_message_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 留言列表
    
    ResultsList   []RefundMessage `json:"results_list,omitempty" xml:"results_list>refund_message,omitempty"`
    
    
}
