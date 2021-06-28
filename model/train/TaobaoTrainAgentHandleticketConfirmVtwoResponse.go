package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商出票中v2--增加鉴权校验 APIResponse
taobao.train.agent.handleticket.confirm.vtwo

代理商出票中
*/
type TaobaoTrainAgentHandleticketConfirmVtwoAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_handleticket_confirm_vtwo_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"