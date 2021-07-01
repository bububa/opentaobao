package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCreativesChangedGetAPIRequest
分页获取修改过的广告创意ID和修改时间 API请求
taobao.simba.creatives.changed.get

分页获取修改过的广告创意ID和修改时间 */
type TaobaoSimbaCreativesChangedGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 返回的每页数据量大小,默认200最大1000
	_pageSize int64
	// 返回的第几页数据，默认为1
	_pageNo int64
	// 得到此时间点之后的数据，不能大于一个月
	_startTime string
}

// New
