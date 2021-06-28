package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除消息topic分组路由 APIResponse
taobao.tmc.topic.group.delete

删除根据topic名称路由消息到不同的分组关系
*/
type TaobaoTmcTopicGroupDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmc_topic_group_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // true
    
    Result   bool `json:"result,omitempty" xml:"