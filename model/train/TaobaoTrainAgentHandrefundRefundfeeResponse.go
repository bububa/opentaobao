package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商手动退款接口 APIResponse
taobao.train.agent.handrefund.refundfee

火车票代理商手动退款接口
*/
type TaobaoTrainAgentHandrefundRefundfeeAPIResponse struct {
    model.CommonResponse
    TaobaoTrainAgentHandrefundRefundfeeResponse
}

type TaobaoTrainAgentHandrefundRefundfeeResponse struct {
    XMLName xml.Name `xml:"train_agent_handrefund_refundfee_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功标记
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 失败code
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 失败文案
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
