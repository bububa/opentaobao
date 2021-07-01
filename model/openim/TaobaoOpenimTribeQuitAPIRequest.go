package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeQuitAPIRequest
OPENIM群成员退出 API请求
taobao.openim.tribe.quit

OPENIM群成员退出 */
type TaobaoOpenimTribeQuitAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
}

// New
