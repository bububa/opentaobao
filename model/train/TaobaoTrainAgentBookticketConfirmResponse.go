package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认出票是否成功 APIResponse
taobao.train.agent.bookticket.confirm

火车票代理商接口——确认出票是否成功
*/
type TaobaoTrainAgentBookticketConfirmAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_bookticket_confirm_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"