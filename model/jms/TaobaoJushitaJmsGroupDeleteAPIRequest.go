package jms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJushitaJmsGroupDeleteAPIRequest
删除ONS分组 API请求
taobao.jushita.jms.group.delete

删除ONS分组 */
type TaobaoJushitaJmsGroupDeleteAPIRequest struct {
	model.Params
	// 分组名称，分组删除后，用户的消息将会存储于默认分组中。警告：由于分组已经删除，用户之前未消费的消息将无法再获取。不能以default开头，default开头为系统默认组。
	_groupName string
	// 用户列表，不传表示删除整个分组，如果用户全部删除后，也会自动删除整个分组
	_nicks []string
	// 用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户
	_userPlatform string
}

// New
