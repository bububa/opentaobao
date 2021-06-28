package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取待改签订单v2--增加鉴权校验 APIResponse
taobao.train.agent.changeorders.get.vtwo

代理商用来获取待改签的订单列表及数量，防止代理商掉单。
*/
type TaobaoTrainAgentChangeordersGetVtwoAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTrainAgentChangeordersGetVtwoResponse `json:"train_agent_changeorders_get_vtwo_response,omitempty"` 
    TaobaoTrainAgentChangeordersGetVtwoResponse
}

/* model for simplify = false
type TaobaoTrainAgentChangeordersGetVtwoResponse struct {

    // 待处理订单总数量
    
    ApplyCount   int64 `json:"apply_count,omitempty"`
    

    // 逗号连接的多个改签单id
    
    ApplyIds   string `json:"apply_ids,omitempty"`
    

}
*/

type TaobaoTrainAgentChangeordersGetVtwoResponse struct {

    // 待处理订单总数量
    ApplyCount   int64 `json:"apply_count,omitempty"`

    // 逗号连接的多个改签单id
    ApplyIds   string `json:"apply_ids,omitempty"`

}
