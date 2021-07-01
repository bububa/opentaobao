package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeModifytribeinfoAPIRequest
OPENIM群信息修改 API请求
taobao.openim.tribe.modifytribeinfo

OPENIM群信息修改 */
type TaobaoOpenimTribeModifytribeinfoAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群名称
	_tribeName string
	// 群公告
	_notice string
	// 群id
	_tribeId int64
}

// New
