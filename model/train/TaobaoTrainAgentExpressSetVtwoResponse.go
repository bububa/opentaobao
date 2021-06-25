package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线下票回填物流信息v2--增加鉴权校验 APIResponse
taobao.train.agent.express.set.vtwo

线下票回填物流信息服务
*/
type TaobaoTrainAgentExpressSetVtwoAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTrainAgentExpressSetVtwoResponse `json:"taobao_train_agent_express_set_vtwo_response,omitempty"`
}

type TaobaoTrainAgentExpressSetVtwoResponse struct {

    // 错误码
    ErrorMsgCode   string `json:"error_msg_code,omitempty"`

    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 扩展参数
    ExtendParams   string `json:"extend_params,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
