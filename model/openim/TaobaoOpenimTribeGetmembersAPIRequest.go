package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeGetmembersAPIRequest
OPENIM群成员获取 API请求
taobao.openim.tribe.getmembers

OPENIM群成员获取 */
type TaobaoOpenimTribeGetmembersAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
}

// New
