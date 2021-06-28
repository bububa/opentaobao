package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取改签单详情v2--增加鉴权校验 APIResponse
taobao.train.agent.change.get.vtwo

卖家获取待处理的改签单详情
*/
type TaobaoTrainAgentChangeGetVtwoAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_change_get_vtwo_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 扩展参数
    
    ExtendParam   string `json:"extend_param,omitempty" xml:"