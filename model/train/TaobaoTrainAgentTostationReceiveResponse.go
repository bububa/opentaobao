package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下票送票至车站代理商确认用户已取票服务 APIResponse
taobao.train.agent.tostation.receive

送票至车站的订单，代理商确认用户已取票
*/
type TaobaoTrainAgentTostationReceiveAPIResponse struct {
    model.CommonResponse
    TaobaoTrainAgentTostationReceiveResponse
}

type TaobaoTrainAgentTostationReceiveResponse struct {
    XMLName xml.Name `xml:"train_agent_tostation_receive_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    ErrorMsgCode   string `json:"error_msg_code,omitempty" xml:"error_msg_code,omitempty"`

    
    // 扩展参数
    
    ExtendParams   string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`

    
    // 错误描述
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
