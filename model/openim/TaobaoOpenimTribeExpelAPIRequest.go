package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeExpelAPIRequest
OPENIM群踢出成员 API请求
taobao.openim.tribe.expel

OPENIM群踢出成员 */
type TaobaoOpenimTribeExpelAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
	// 用户信息
	_member *OpenImUser
}

// New
