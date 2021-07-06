package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelTradesSearchAPIRequest 飞猪度假-订单列表搜索接口 API请求
// alitrip.travel.trades.search
//
// 订单列表搜索接口：以订单创建、结束时间、订单状态为搜索条件，搜索过滤出满足条件的卖家订单列表。
type AlitripTravelTradesSearchAPIRequest struct {
	model.Params
	// 订单创建 结束时间
	_endCreatedTime string
	// 订单创建 开始时间
	_startCreatedTime string
	// 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
	_pageSize int64
	// 订单状态 过滤。1-等待买家付款，2-等待卖家发货（买家已付款），3-等待买家确认收货，4-交易关闭（买家发起的退款），6-交易成功，8-交易关闭（订单超时 自动关单）
	_orderStatus int64
	// 当前页
	_currentPage int64
	// 类目筛选, 1、旅行购，旅行购定制专用字段，表示搜索旅行购订单。
	_category int64
}

// NewAlitripTravelTradesSearchRequest 初始化AlitripTravelTradesSearchAPIRequest对象
func NewAlitripTravelTradesSearchRequest() *AlitripTravelTradesSearchAPIRequest {
	return &AlitripTravelTradesSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelTradesSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.trades.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelTradesSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetEndCreatedTime is EndCreatedTime Setter
// 订单创建 结束时间
func (r *AlitripTravelTradesSearchAPIRequest) SetEndCreatedTime(_endCreatedTime string) error {
	r._endCreatedTime = _endCreatedTime
	r.Set("end_created_time", _endCreatedTime)
	return nil
}

// GetEndCreatedTime EndCreatedTime Getter
func (r AlitripTravelTradesSearchAPIRequest) GetEndCreatedTime() string {
	return r._endCreatedTime
}

// SetStartCreatedTime is StartCreatedTime Setter
// 订单创建 开始时间
func (r *AlitripTravelTradesSearchAPIRequest) SetStartCreatedTime(_startCreatedTime string) error {
	r._startCreatedTime = _startCreatedTime
	r.Set("start_created_time", _startCreatedTime)
	return nil
}

// GetStartCreatedTime StartCreatedTime Getter
func (r AlitripTravelTradesSearchAPIRequest) GetStartCreatedTime() string {
	return r._startCreatedTime
}

// SetPageSize is PageSize Setter
// 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
func (r *AlitripTravelTradesSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlitripTravelTradesSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetOrderStatus is OrderStatus Setter
// 订单状态 过滤。1-等待买家付款，2-等待卖家发货（买家已付款），3-等待买家确认收货，4-交易关闭（买家发起的退款），6-交易成功，8-交易关闭（订单超时 自动关单）
func (r *AlitripTravelTradesSearchAPIRequest) SetOrderStatus(_orderStatus int64) error {
	r._orderStatus = _orderStatus
	r.Set("order_status", _orderStatus)
	return nil
}

// GetOrderStatus OrderStatus Getter
func (r AlitripTravelTradesSearchAPIRequest) GetOrderStatus() int64 {
	return r._orderStatus
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *AlitripTravelTradesSearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlitripTravelTradesSearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetCategory is Category Setter
// 类目筛选, 1、旅行购，旅行购定制专用字段，表示搜索旅行购订单。
func (r *AlitripTravelTradesSearchAPIRequest) SetCategory(_category int64) error {
	r._category = _category
	r.Set("category", _category)
	return nil
}

// GetCategory Category Getter
func (r AlitripTravelTradesSearchAPIRequest) GetCategory() int64 {
	return r._category
}
