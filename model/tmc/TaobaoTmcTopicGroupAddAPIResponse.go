package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
topic分组路由 API返回值 
taobao.tmc.topic.group.add

根据topic名称路由消息到不同的分组。（前提：发送方未指定分组名）
如果是需要授权的消息，分组路由先判断用户分组路由(使用taobao.tmc.group.add添加的路由)，用户分组路由不存在时，才会判断topic分组路由
*/
type TaobaoTmcTopicGroupAddAPIResponse struct {
    model.CommonResponse
    TaobaoTmcTopicGroupAddAPIResponseModel
}

// topic分组路由 成功返回结果
type TaobaoTmcTopicGroupAddAPIResponseModel struct {
    XMLName xml.Name `xml:"tmc_topic_group_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // true
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
