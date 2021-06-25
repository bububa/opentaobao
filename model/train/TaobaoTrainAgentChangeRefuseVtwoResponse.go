package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
代理商拒绝改签v2--增加鉴权校验 APIResponse
taobao.train.agent.change.refuse.vtwo

代理商拒绝火车票改签服务
*/
type TaobaoTrainAgentChangeRefuseVtwoAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTrainAgentChangeRefuseVtwoResponse `json:"taobao_train_agent_change_refuse_vtwo_response,omitempty"`
}

type TaobaoTrainAgentChangeRefuseVtwoResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
