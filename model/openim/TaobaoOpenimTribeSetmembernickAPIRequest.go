package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeSetmembernickAPIRequest
设置群成员昵称 API请求
taobao.openim.tribe.setmembernick

设置群成员昵称，存在如下两种场景
1 群主或管理员设置群成员昵称，该操作有权限控制。只针对普通群的群主和管理员开发此功能；讨论组群主不支持此设置操作
2 群成员设置自己的昵称，该功能对群所有成员开放 */
type TaobaoOpenimTribeSetmembernickAPIRequest struct {
	model.Params
	// 发起设置昵称的操作者，如果是设置其他成员的昵称，只有普通组的群主和管理员有权限
	_user *User
	// 群id
	_tribeId int64
	// 被设置昵称的群成员
	_member *User
	// 设置的昵称
	_nick string
}

// New
