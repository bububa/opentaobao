package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除消息topic分组路由 APIResponse
taobao.tmc.topic.group.delete

删除根据topic名称路由消息到不同的分组关系
*/
type TaobaoTmcTopicGroupDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTmcTopicGroupDeleteResponse `json:"taobao_tmc_topic_group_delete_response,omitempty"`
}

type TaobaoTmcTopicGroupDeleteResponse struct {

    // true
    Result   bool `json:"result,omitempty"`

}
