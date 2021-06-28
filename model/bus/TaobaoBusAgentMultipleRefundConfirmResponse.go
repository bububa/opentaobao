package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
综合交通多次退款接口 APIResponse
taobao.bus.agent.multiple.refund.confirm

此接口支持多次按照单客进行多次退款操作，只进行退款操作。
*/
type TaobaoBusAgentMultipleRefundConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoBusAgentMultipleRefundConfirmResponse
}

type TaobaoBusAgentMultipleRefundConfirmResponse struct {
    XMLName xml.Name `xml:"bus_agent_multiple_refund_confirm_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 失败错误信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 退款成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
