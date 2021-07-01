package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeGetalltribesAPIRequest
获取用户群列表 API请求
taobao.openim.tribe.getalltribes

OPENIM群服务获取用户群列表 */
type TaobaoOpenimTribeGetalltribesAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群类型
	_tribeTypes []int64
}

// New
