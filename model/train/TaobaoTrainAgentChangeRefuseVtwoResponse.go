package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商拒绝改签v2--增加鉴权校验 APIResponse
taobao.train.agent.change.refuse.vtwo

代理商拒绝火车票改签服务
*/
type TaobaoTrainAgentChangeRefuseVtwoAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_change_refuse_vtwo_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"