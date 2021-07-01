package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribelogsGetAPIRequest
openim 群聊天记录导出接口 API请求
taobao.openim.tribelogs.get

获取openim账号的群聊天记录 */
type TaobaoOpenimTribelogsGetAPIRequest struct {
	model.Params
	// 群号
	_tribeId string
	// 查询起始时间，UTC秒数。必须在一个月内。
	_begin int64
	// 查询结束时间，UTC秒数。必须大于起始时间并小于当前时间
	_end int64
	// 查询条数
	_count int64
	// 迭代key
	_next string
}

// New
