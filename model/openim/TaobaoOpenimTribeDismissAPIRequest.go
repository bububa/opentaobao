package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeDismissAPIRequest
OPENIM群解散 API请求
taobao.openim.tribe.dismiss

OPENIM群解散 */
type TaobaoOpenimTribeDismissAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
}

// New
