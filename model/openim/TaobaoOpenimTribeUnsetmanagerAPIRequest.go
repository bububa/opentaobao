package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeUnsetmanagerAPIRequest
OPENIM群取消管理员 API请求
taobao.openim.tribe.unsetmanager

OPENIM群取消管理员 */
type TaobaoOpenimTribeUnsetmanagerAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tid int64
	// 用户信息
	_member *OpenImUser
}

// New
