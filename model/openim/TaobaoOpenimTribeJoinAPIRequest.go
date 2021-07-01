package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeJoinAPIRequest
OPENIM群主动加入 API请求
taobao.openim.tribe.join

OPENIM群主动加入 */
type TaobaoOpenimTribeJoinAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
}

// New
