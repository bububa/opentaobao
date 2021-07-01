package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcTopicGroupAddAPIRequest
topic分组路由 API请求
taobao.tmc.topic.group.add

根据topic名称路由消息到不同的分组。（前提：发送方未指定分组名）
如果是需要授权的消息，分组路由先判断用户分组路由(使用taobao.tmc.group.add添加的路由)，用户分组路由不存在时，才会判断topic分组路由 */
type TaobaoTmcTopicGroupAddAPIRequest struct {
	model.Params
	// 消息分组名，如果不存在，会自动创建
	_groupName string
	// 消息topic名称，多个以逗号(,)分割
	_topics []string
}

// New
