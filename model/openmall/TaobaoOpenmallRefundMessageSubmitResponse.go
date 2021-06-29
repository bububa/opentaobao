package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交退款单留言 API返回值 
taobao.openmall.refund.message.submit

OpenMall业务提交退款单留言
*/
type TaobaoOpenmallRefundMessageSubmitAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallRefundMessageSubmitResponse
}

// 提交退款单留言 成功返回结果
type TaobaoOpenmallRefundMessageSubmitResponse struct {
    XMLName xml.Name `xml:"openmall_refund_message_submit_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 提交结果
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
