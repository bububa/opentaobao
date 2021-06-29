package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认改签占座是否成功 APIResponse
taobao.train.agent.change.holdseat.confirm

火车票代理商接口——确认改签占座是否成功
*/
type TaobaoTrainAgentChangeHoldseatConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoTrainAgentChangeHoldseatConfirmResponse
}

type TaobaoTrainAgentChangeHoldseatConfirmResponse struct {
    XMLName xml.Name `xml:"train_agent_change_holdseat_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功标记
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // errorCode
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // errorMsg
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
