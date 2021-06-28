package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取待改签订单 APIResponse
taobao.train.agent.changeorders.get

代理商用来获取待改签的订单列表及数量，防止代理商掉单。
*/
type TaobaoTrainAgentChangeordersGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTrainAgentChangeordersGetResponse `json:"train_agent_changeorders_get_response,omitempty"` 
    TaobaoTrainAgentChangeordersGetResponse
}

/* model for simplify = false
type TaobaoTrainAgentChangeordersGetResponse struct {

    // 待处理订单总数量
    
    ApplyCount   int64 `json:"apply_count,omitempty"`
    

    // 逗号连接的多个改签单id
    
    ApplyIds   string `json:"apply_ids,omitempty"`
    

}
*/

type TaobaoTrainAgentChangeordersGetResponse struct {

    // 待处理订单总数量
    ApplyCount   int64 `json:"apply_count,omitempty"`

    // 逗号连接的多个改签单id
    ApplyIds   string `json:"apply_ids,omitempty"`

}
