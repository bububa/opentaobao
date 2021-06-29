package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】查询辅营订单列表 API请求
alitrip.tripvp.agent.order.search

【国际机票】查询辅营订单列表
*/
type AlitripTripvpAgentOrderSearchRequest struct {
    model.Params
    // 代理商ID
    _agentId   int64
    // 辅营创建开始时间
    _beginTime   string
    // 当前页码
    _currentPage   int64
    // 辅营创建结束时间
    _endTime   string
    // 订单状态，1-待支付 2-支付成功 3-	辅营出货成功 4-订单取消
    _orderStatus   int64
    // 分页行数
    _pageSize   int64
}

// 初始化AlitripTripvpAgentOrderSearchRequest对象
func NewAlitripTripvpAgentOrderSearchRequest() *AlitripTripvpAgentOrderSearchRequest{
    return &AlitripTripvpAgentOrderSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTripvpAgentOrderSearchRequest) GetApiMethodName() string {
    return "alitrip.tripvp.agent.order.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTripvpAgentOrderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 代理商ID
func (r *AlitripTripvpAgentOrderSearchRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r AlitripTripvpAgentOrderSearchRequest) GetAgentId() int64 {
    return r._agentId
}
// BeginTime Setter
// 辅营创建开始时间
func (r *AlitripTripvpAgentOrderSearchRequest) SetBeginTime(_beginTime string) error {
    r._beginTime = _beginTime
    r.Set("begin_time", _beginTime)
    return nil
}

// BeginTime Getter
func (r AlitripTripvpAgentOrderSearchRequest) GetBeginTime() string {
    return r._beginTime
}
// CurrentPage Setter
// 当前页码
func (r *AlitripTripvpAgentOrderSearchRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlitripTripvpAgentOrderSearchRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// EndTime Setter
// 辅营创建结束时间
func (r *AlitripTripvpAgentOrderSearchRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r AlitripTripvpAgentOrderSearchRequest) GetEndTime() string {
    return r._endTime
}
// OrderStatus Setter
// 订单状态，1-待支付 2-支付成功 3-	辅营出货成功 4-订单取消
func (r *AlitripTripvpAgentOrderSearchRequest) SetOrderStatus(_orderStatus int64) error {
    r._orderStatus = _orderStatus
    r.Set("order_status", _orderStatus)
    return nil
}

// OrderStatus Getter
func (r AlitripTripvpAgentOrderSearchRequest) GetOrderStatus() int64 {
    return r._orderStatus
}
// PageSize Setter
// 分页行数
func (r *AlitripTripvpAgentOrderSearchRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlitripTripvpAgentOrderSearchRequest) GetPageSize() int64 {
    return r._pageSize
}