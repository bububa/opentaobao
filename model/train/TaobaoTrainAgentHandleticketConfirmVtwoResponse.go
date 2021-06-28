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
    TaobaoTrainAgentHandleticketConfirmVtwoResponse
}

type TaobaoTrainAgentHandleticketConfirmVtwoResponse struct {
    XMLName xml.Name `xml:"train_agent_handleticket_confirm_vtwo_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 错误码
    
    TrainErrorCode   string `json:"train_error_code,omitempty" xml:"train_error_code,omitempty"`

    
    // 错误信息
    
    TrainErrorMsg   string `json:"train_error_msg,omitempty" xml:"train_error_msg,omitempty"`

    
    // 暂无
    
    ExtendParams   string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`

    
}
