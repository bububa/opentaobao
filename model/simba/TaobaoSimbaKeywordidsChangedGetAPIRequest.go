package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordidsChangedGetAPIRequest
获取修改的词ID API请求
taobao.simba.keywordids.changed.get

获取修改的词ID */
type TaobaoSimbaKeywordidsChangedGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 得到此时间点之后的数据，不能大于一个月
	_startTime string
	// 返回的每页数据量大小,默认200最大1000
	_pageSize int64
	// 返回的第几页数据，默认为1
	_pageNo int64
}

// New
