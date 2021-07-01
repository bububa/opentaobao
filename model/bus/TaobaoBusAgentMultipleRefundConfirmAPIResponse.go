package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
综合交通多次退款接口 API返回值 
taobao.bus.agent.multiple.refund.confirm

此接口支持多次按照单客进行多次退款操作，只进行退款操作。
*/
type TaobaoBusAgentMultipleRefundConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoBusAgentMultipleRefundConfirmAPIResponseModel
}

// 综合交通多次退款接口 成功返回结果
type TaobaoBusAgentMultipleRefundConfirmAPIResponseModel struct {
    XMLName xml.Name `xml:"bus_agent_multiple_refund_confirm_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 失败错误信息
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // 退款成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
