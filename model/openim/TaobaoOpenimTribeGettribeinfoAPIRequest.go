package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeGettribeinfoAPIRequest
获取群信息 API请求
taobao.openim.tribe.gettribeinfo

获取群信息 */
type TaobaoOpenimTribeGettribeinfoAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群ID
	_tribeId int64
}

// New
