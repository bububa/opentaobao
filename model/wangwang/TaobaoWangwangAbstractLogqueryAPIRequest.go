package wangwang

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWangwangAbstractLogqueryAPIRequest
模糊聊天记录查询 API请求
taobao.wangwang.abstract.logquery

模糊聊天记录查询 */
type TaobaoWangwangAbstractLogqueryAPIRequest struct {
	model.Params
	// 买家id，有cntaobao前缀
	_toId string
	// 卖家id，有cntaobao前缀
	_fromId string
	// 获取记录条数，默认值是1000
	_count int64
	// 设置了这个值，那么聊天记录就从这个点开始查询
	_nextKey string
	// 传入参数的字符集
	_charset string
	// utc
	_startDate int64
	// utc
	_endDate int64
}

// New
