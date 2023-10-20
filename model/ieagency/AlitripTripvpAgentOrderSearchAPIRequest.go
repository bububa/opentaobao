package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptripvpagentordersearchAPIRequest 【国际机票】查询辅营订单列表 API请求
// alitrip.tripvp.agent.order.search
//
// 【国际机票】查询辅营订单列表
type AlitriptripvpagentordersearchAPIRequest struct {
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

// NewAlitriptripvpagentordersearchRequest 初始化AlitriptripvpagentordersearchAPIRequest对象
func NewAlitriptripvpagentordersearchRequest() *AlitriptripvpagentordersearchAPIRequest {
	return &AlitriptripvpagentordersearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptripvpagentordersearchAPIRequest) GetApiMethodName() string {
	return "alitrip.tripvp.agent.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptripvpagentordersearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptripvpagentordersearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBeginTime is BeginTime Setter
// 辅营创建开始时间
func (r *AlitriptripvpagentordersearchAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// GetBeginTime BeginTime Getter
func (r AlitriptripvpagentordersearchAPIRequest) GetBeginTime() string {
	return r._beginTime
}

// SetEndTime is EndTime Setter
// 辅营创建结束时间
func (r *AlitriptripvpagentordersearchAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlitriptripvpagentordersearchAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetAgentId is AgentId Setter
// 代理商ID
func (r *AlitriptripvpagentordersearchAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlitriptripvpagentordersearchAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetCurrentPage is CurrentPage Setter
// 当前页码
func (r *AlitriptripvpagentordersearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlitriptripvpagentordersearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetOrderStatus is OrderStatus Setter
// 订单状态，1-待支付 2-支付成功 3-	辅营出货成功 4-订单取消
func (r *AlitriptripvpagentordersearchAPIRequest) SetOrderStatus(_orderStatus int64) error {
	r._orderStatus = _orderStatus
	r.Set("order_status", _orderStatus)
	return nil
}

// GetOrderStatus OrderStatus Getter
func (r AlitriptripvpagentordersearchAPIRequest) GetOrderStatus() int64 {
	return r._orderStatus
}

// SetPageSize is PageSize Setter
// 分页行数
func (r *AlitriptripvpagentordersearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlitriptripvpagentordersearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
