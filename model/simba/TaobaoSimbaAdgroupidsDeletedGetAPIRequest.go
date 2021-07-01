package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaAdgroupidsDeletedGetAPIRequest
获取删除的推广组ID API请求
taobao.simba.adgroupids.deleted.get

获取删除的推广组ID */
type TaobaoSimbaAdgroupidsDeletedGetAPIRequest struct {
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
