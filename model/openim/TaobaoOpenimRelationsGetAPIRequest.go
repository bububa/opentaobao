package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimRelationsGetAPIRequest
获取openim账号的聊天关系 API请求
taobao.openim.relations.get

获取openim账号的聊天关系 */
type TaobaoOpenimRelationsGetAPIRequest struct {
	model.Params
	// 查询起始日期。格式yyyyMMdd。不得早于一个月
	_begDate string
	// 查询结束日期。格式yyyyMMdd。不得早于一个月
	_endDate string
	// 用户信息
	_user *OpenImUser
}

// New
