package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
topic分组路由 APIResponse
taobao.tmc.topic.group.add

根据topic名称路由消息到不同的分组。（前提：发送方未指定分组名）
如果是需要授权的消息，分组路由先判断用户分组路由(使用taobao.tmc.group.add添加的路由)，用户分组路由不存在时，才会判断topic分组路由
*/
type TaobaoTmcTopicGroupAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTmcTopicGroupAddResponse `json:"tmc_topic_group_add_response,omitempty"` 
    TaobaoTmcTopicGroupAddResponse
}

/* model for simplify = false
type TaobaoTmcTopicGroupAddResponse struct {

    // true
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type TaobaoTmcTopicGroupAddResponse struct {

    // true
    Result   bool `json:"result,omitempty"`

}
