package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取待改签订单 APIResponse
taobao.train.agent.changeorders.get

代理商用来获取待改签的订单列表及数量，防止代理商掉单。
*/
type TaobaoTrainAgentChangeordersGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_changeorders_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 待处理订单总数量
    
    ApplyCount   int64 `json:"apply_count,omitempty" xml:"