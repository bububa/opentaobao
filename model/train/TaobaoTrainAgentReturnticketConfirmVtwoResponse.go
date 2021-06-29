package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退票通知 APIResponse
taobao.train.agent.returnticket.confirm.vtwo

火车票代理商接口——退票通知回调
*/
type TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse struct {
    model.CommonResponse
    TaobaoTrainAgentReturnticketConfirmVtwoResponse
}

type TaobaoTrainAgentReturnticketConfirmVtwoResponse struct {
    XMLName xml.Name `xml:"train_agent_returnticket_confirm_vtwo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // resultCode
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // resultMsg
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
