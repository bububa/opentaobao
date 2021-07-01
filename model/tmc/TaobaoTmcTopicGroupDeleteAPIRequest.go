package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcTopicGroupDeleteAPIRequest
删除消息topic分组路由 API请求
taobao.tmc.topic.group.delete

删除根据topic名称路由消息到不同的分组关系 */
type TaobaoTmcTopicGroupDeleteAPIRequest struct {
	model.Params
	// 消息分组名
	_groupName string
	// 消息topic名称，多个以逗号(,)分割
	_topics []string
	// 消息分组Id，一般不用填写，如果分组已经被删除，则根据问题排查工具返回的ID删除路由关系
	_groupId int64
}

// New
