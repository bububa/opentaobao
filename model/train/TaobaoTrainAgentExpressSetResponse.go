package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线下票回填物流信息 APIResponse
taobao.train.agent.express.set

线下票回填物流信息服务
*/
type TaobaoTrainAgentExpressSetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTrainAgentExpressSetResponse `json:"train_agent_express_set_response,omitempty"` 
    TaobaoTrainAgentExpressSetResponse
}

/* model for simplify = false
type TaobaoTrainAgentExpressSetResponse struct {

    // 错误码
    
    ErrorMsgCode   string `json:"error_msg_code,omitempty"`
    

    // 错误描述
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 扩展参数
    
    ExtendParams   string `json:"extend_params,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoTrainAgentExpressSetResponse struct {

    // 错误码
    ErrorMsgCode   string `json:"error_msg_code,omitempty"`

    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 扩展参数
    ExtendParams   string `json:"extend_params,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
