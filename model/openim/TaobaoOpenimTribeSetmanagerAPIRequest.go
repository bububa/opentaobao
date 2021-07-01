package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeSetmanagerAPIRequest
OPENIM群设置管理员 API请求
taobao.openim.tribe.setmanager

OPENIM群设置管理员 */
type TaobaoOpenimTribeSetmanagerAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tid int64
	// 用户信息
	_member *OpenImUser
}

// New
