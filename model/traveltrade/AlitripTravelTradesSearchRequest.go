package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单列表搜索接口 API请求
alitrip.travel.trades.search

订单列表搜索接口：以订单创建、结束时间、订单状态为搜索条件，搜索过滤出满足条件的卖家订单列表。
*/
type AlitripTravelTradesSearchRequest struct {
    model.Params
    // 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
    _pageSize   int64
    // 订单创建 结束时间
    _endCreatedTime   string
    // 订单状态 过滤。1-等待买家付款，2-等待卖家发货（买家已付款），3-等待买家确认收货，4-交易关闭（买家发起的退款），6-交易成功，8-交易关闭（订单超时 自动关单）
    _orderStatus   int64
    // 当前页
    _currentPage   int64
    // 订单创建 开始时间
    _startCreatedTime   string
    // 类目筛选, 1、旅行购，旅行购定制专用字段，表示搜索旅行购订单。
    _category   int64
}

// 初始化AlitripTravelTradesSearchRequest对象
func NewAlitripTravelTradesSearchRequest() *AlitripTravelTradesSearchRequest{
    return &AlitripTravelTradesSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelTradesSearchRequest) GetApiMethodName() string {
    return "alitrip.travel.trades.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelTradesSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
func (r *AlitripTravelTradesSearchRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlitripTravelTradesSearchRequest) GetPageSize() int64 {
    return r._pageSize
}
// EndCreatedTime Setter
// 订单创建 结束时间
func (r *AlitripTravelTradesSearchRequest) SetEndCreatedTime(_endCreatedTime string) error {
    r._endCreatedTime = _endCreatedTime
    r.Set("end_created_time", _endCreatedTime)
    return nil
}

// EndCreatedTime Getter
func (r AlitripTravelTradesSearchRequest) GetEndCreatedTime() string {
    return r._endCreatedTime
}
// OrderStatus Setter
// 订单状态 过滤。1-等待买家付款，2-等待卖家发货（买家已付款），3-等待买家确认收货，4-交易关闭（买家发起的退款），6-交易成功，8-交易关闭（订单超时 自动关单）
func (r *AlitripTravelTradesSearchRequest) SetOrderStatus(_orderStatus int64) error {
    r._orderStatus = _orderStatus
    r.Set("order_status", _orderStatus)
    return nil
}

// OrderStatus Getter
func (r AlitripTravelTradesSearchRequest) GetOrderStatus() int64 {
    return r._orderStatus
}
// CurrentPage Setter
// 当前页
func (r *AlitripTravelTradesSearchRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlitripTravelTradesSearchRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// StartCreatedTime Setter
// 订单创建 开始时间
func (r *AlitripTravelTradesSearchRequest) SetStartCreatedTime(_startCreatedTime string) error {
    r._startCreatedTime = _startCreatedTime
    r.Set("start_created_time", _startCreatedTime)
    return nil
}

// StartCreatedTime Getter
func (r AlitripTravelTradesSearchRequest) GetStartCreatedTime() string {
    return r._startCreatedTime
}
// Category Setter
// 类目筛选, 1、旅行购，旅行购定制专用字段，表示搜索旅行购订单。
func (r *AlitripTravelTradesSearchRequest) SetCategory(_category int64) error {
    r._category = _category
    r.Set("category", _category)
    return nil
}

// Category Getter
func (r AlitripTravelTradesSearchRequest) GetCategory() int64 {
    return r._category
}
