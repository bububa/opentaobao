package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimChatlogsGetAPIRequest
openim聊天记录查询接口 API请求
taobao.openim.chatlogs.get

查询openim账号聊天记录 */
type TaobaoOpenimChatlogsGetAPIRequest struct {
	model.Params
	// 用户1信息
	_user1 *OpenImUser
	// 用户2信息
	_user2 *OpenImUser
	// 查询开始时间（UTC时间）
	_begin int64
	// 查询结束时间（UTC时间）
	_end int64
	// 查询条数
	_count int64
	// 迭代key
	_nextKey string
}

// New
