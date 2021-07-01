package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimAppChatlogsGetAPIRequest
openim应用聊天记录查询 API请求
taobao.openim.app.chatlogs.get

查询openim应用的聊天记录 */
type TaobaoOpenimAppChatlogsGetAPIRequest struct {
	model.Params
	// 查询结束时间。UTC时间。精度到秒
	_beg int64
	// 查询结束时间。UTC时间。精度到秒
	_end int64
	// 查询最大条数
	_count int64
	// 迭代key
	_next string
}

// New
