package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptraveltradessearchAPIRequest 飞猪度假-订单列表搜索接口 API请求
// alitrip.travel.trades.search
//
// 订单列表搜索接口：以订单创建、结束时间、订单状态为搜索条件，搜索过滤出满足条件的卖家订单列表。
type AlitriptraveltradessearchAPIRequest struct {
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
	// 商品ID
	_itemId int64
}

// NewAlitriptraveltradessearchRequest 初始化AlitriptraveltradessearchAPIRequest对象
func NewAlitriptraveltradessearchRequest() *AlitriptraveltradessearchAPIRequest {
	return &AlitriptraveltradessearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptraveltradessearchAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.trades.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptraveltradessearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptraveltradessearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndCreatedTime is EndCreatedTime Setter
// 订单创建 结束时间
func (r *AlitriptraveltradessearchAPIRequest) SetEndCreatedTime(_endCreatedTime string) error {
	r._endCreatedTime = _endCreatedTime
	r.Set("end_created_time", _endCreatedTime)
	return nil
}

// GetEndCreatedTime EndCreatedTime Getter
func (r AlitriptraveltradessearchAPIRequest) GetEndCreatedTime() string {
	return r._endCreatedTime
}

// SetStartCreatedTime is StartCreatedTime Setter
// 订单创建 开始时间
func (r *AlitriptraveltradessearchAPIRequest) SetStartCreatedTime(_startCreatedTime string) error {
	r._startCreatedTime = _startCreatedTime
	r.Set("start_created_time", _startCreatedTime)
	return nil
}

// GetStartCreatedTime StartCreatedTime Getter
func (r AlitriptraveltradessearchAPIRequest) GetStartCreatedTime() string {
	return r._startCreatedTime
}

// SetPageSize is PageSize Setter
// 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
func (r *AlitriptraveltradessearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlitriptraveltradessearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetOrderStatus is OrderStatus Setter
// 订单状态 过滤。1-等待买家付款，2-等待卖家发货（买家已付款），3-等待买家确认收货，4-交易关闭（买家发起的退款），6-交易成功，8-交易关闭（订单超时 自动关单）
func (r *AlitriptraveltradessearchAPIRequest) SetOrderStatus(_orderStatus int64) error {
	r._orderStatus = _orderStatus
	r.Set("order_status", _orderStatus)
	return nil
}

// GetOrderStatus OrderStatus Getter
func (r AlitriptraveltradessearchAPIRequest) GetOrderStatus() int64 {
	return r._orderStatus
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *AlitriptraveltradessearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlitriptraveltradessearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetCategory is Category Setter
// 类目筛选, 1、旅行购，旅行购定制专用字段，表示搜索旅行购订单。
func (r *AlitriptraveltradessearchAPIRequest) SetCategory(_category int64) error {
	r._category = _category
	r.Set("category", _category)
	return nil
}

// GetCategory Category Getter
func (r AlitriptraveltradessearchAPIRequest) GetCategory() int64 {
	return r._category
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlitriptraveltradessearchAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitriptraveltradessearchAPIRequest) GetItemId() int64 {
	return r._itemId
}
