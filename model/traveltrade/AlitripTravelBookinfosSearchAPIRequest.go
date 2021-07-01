package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelBookinfosSearchAPIRequest
飞猪度假-订单预定信息列表搜索接口 API请求
alitrip.travel.bookinfos.search

查询订单预定信息列表 */
type AlitripTravelBookinfosSearchAPIRequest struct {
	model.Params
	// 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
	_pageSize int64
	// 当前页
	_currentPage int64
	// 申请时间_结束，精确到分钟
	_applyTimeEnd string
	// 申请时间_开始，精确到分钟
	_applyTimeStart string
}

// New
