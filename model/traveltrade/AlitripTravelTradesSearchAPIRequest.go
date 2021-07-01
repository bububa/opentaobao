package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelTradesSearchAPIRequest
飞猪度假-订单列表搜索接口 API请求
alitrip.travel.trades.search

订单列表搜索接口：以订单创建、结束时间、订单状态为搜索条件，搜索过滤出满足条件的卖家订单列表。 */
type AlitripTravelTradesSearchAPIRequest struct {
	model.Params
	// 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
	_pageSize int64
	// 订单创建 结束时间
	_endCreatedTime string
	// 订单状态 过滤。1-等待买家付款，2-等待卖家发货（买家已付款），3-等待买家确认收货，4-交易关闭（买家发起的退款），6-交易成功，8-交易关闭（订单超时 自动关单）
	_orderStatus int64
	// 当前页
	_currentPage int64
	// 订单创建 开始时间
	_startCreatedTime string
	// 类目筛选, 1、旅行购，旅行购定制专用字段，表示搜索旅行购订单。
	_category int64
}

// New
