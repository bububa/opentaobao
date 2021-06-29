package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认占座是否成功 APIResponse
taobao.train.agent.holdseat.confirm

火车票代理商接口——确认占座是否成功
*/
type TaobaoTrainAgentHoldseatConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoTrainAgentHoldseatConfirmResponse
}

type TaobaoTrainAgentHoldseatConfirmResponse struct {
    XMLName xml.Name `xml:"train_agent_holdseat_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // resultMsg
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // resultCode
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
