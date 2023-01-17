package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTripvpAgentOrderSearchAPIRequest 【国际机票】查询辅营订单列表 API请求
// alitrip.tripvp.agent.order.search
//
// 【国际机票】查询辅营订单列表
type AlitripTripvpAgentOrderSearchAPIRequest struct {
	model.Params
	// 辅营创建开始时间
	_beginTime string
	// 辅营创建结束时间
	_endTime string
	// 代理商ID
	_agentId int64
	// 当前页码
	_currentPage int64
	// 订单状态，1-待支付 2-支付成功 3-	辅营出货成功 4-订单取消
	_orderStatus int64
	// 分页行数
	_pageSize int64
}

// NewAlitripTripvpAgentOrderSearchRequest 初始化AlitripTripvpAgentOrderSearchAPIRequest对象
func NewAlitripTripvpAgentOrderSearchRequest() *AlitripTripvpAgentOrderSearchAPIRequest {
	return &AlitripTripvpAgentOrderSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTripvpAgentOrderSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.tripvp.agent.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTripvpAgentOrderSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTripvpAgentOrderSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBeginTime is BeginTime Setter
// 辅营创建开始时间
func (r *AlitripTripvpAgentOrderSearchAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// GetBeginTime BeginTime Getter
func (r AlitripTripvpAgentOrderSearchAPIRequest) GetBeginTime() string {
	return r._beginTime
}

// SetEndTime is EndTime Setter
// 辅营创建结束时间
func (r *AlitripTripvpAgentOrderSearchAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlitripTripvpAgentOrderSearchAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetAgentId is AgentId Setter
// 代理商ID
func (r *AlitripTripvpAgentOrderSearchAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlitripTripvpAgentOrderSearchAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetCurrentPage is CurrentPage Setter
// 当前页码
func (r *AlitripTripvpAgentOrderSearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlitripTripvpAgentOrderSearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetOrderStatus is OrderStatus Setter
// 订单状态，1-待支付 2-支付成功 3-	辅营出货成功 4-订单取消
func (r *AlitripTripvpAgentOrderSearchAPIRequest) SetOrderStatus(_orderStatus int64) error {
	r._orderStatus = _orderStatus
	r.Set("order_status", _orderStatus)
	return nil
}

// GetOrderStatus OrderStatus Getter
func (r AlitripTripvpAgentOrderSearchAPIRequest) GetOrderStatus() int64 {
	return r._orderStatus
}

// SetPageSize is PageSize Setter
// 分页行数
func (r *AlitripTripvpAgentOrderSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlitripTripvpAgentOrderSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
