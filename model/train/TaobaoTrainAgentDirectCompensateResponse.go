package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——订单关闭实际出票成功审计接口 APIResponse
taobao.train.agent.direct.compensate

代购直连订单平台关单但是代理商出票成功补偿接口
*/
type TaobaoTrainAgentDirectCompensateAPIResponse struct {
    model.CommonResponse
    TaobaoTrainAgentDirectCompensateResponse
}

type TaobaoTrainAgentDirectCompensateResponse struct {
    XMLName xml.Name `xml:"train_agent_direct_compensate_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // resultCode
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // resultMsg
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
