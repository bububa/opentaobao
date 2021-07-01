package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeInviteAPIRequest
OPENIM群邀请加入 API请求
taobao.openim.tribe.invite

OPENIM群邀请加入接口 */
type TaobaoOpenimTribeInviteAPIRequest struct {
	model.Params
	// 群id
	_tribeId int64
	// 用户信息
	_members []OpenImUser
	// 用户信息
	_user *OpenImUser
}

// New
