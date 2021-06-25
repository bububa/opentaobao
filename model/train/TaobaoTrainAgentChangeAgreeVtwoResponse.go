package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
代理商同意改签v2--增加鉴权校验 APIResponse
taobao.train.agent.change.agree.vtwo

代理商同意改签接口服务
*/
type TaobaoTrainAgentChangeAgreeVtwoAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTrainAgentChangeAgreeVtwoResponse `json:"taobao_train_agent_change_agree_vtwo_response,omitempty"`
}

type TaobaoTrainAgentChangeAgreeVtwoResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
