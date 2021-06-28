package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商同意改签v2--增加鉴权校验 APIResponse
taobao.train.agent.change.agree.vtwo

代理商同意改签接口服务
*/
type TaobaoTrainAgentChangeAgreeVtwoAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_change_agree_vtwo_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"