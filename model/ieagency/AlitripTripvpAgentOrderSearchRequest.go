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
    agentId   int64
    // 辅营创建开始时间
    beginTime   string
    // 当前页码
    currentPage   int64
    // 辅营创建结束时间
    endTime   string
    // 订单状态，1-待支付 2-支付成功 3-	辅营出货成功 4-订单取消
    orderStatus   int64
    // 分页行数
    pageSize   int64
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
func (r *AlitripTripvpAgentOrderSearchRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r AlitripTripvpAgentOrderSearchRequest) GetAgentId() int64 {
    return r.agentId
}
// BeginTime Setter
// 辅营创建开始时间
func (r *AlitripTripvpAgentOrderSearchRequest) SetBeginTime(beginTime string) error {
    r.beginTime = beginTime
    r.Set("begin_time", beginTime)
    return nil
}

// BeginTime Getter
func (r AlitripTripvpAgentOrderSearchRequest) GetBeginTime() string {
    return r.beginTime
}
// CurrentPage Setter
// 当前页码
func (r *AlitripTripvpAgentOrderSearchRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r AlitripTripvpAgentOrderSearchRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// EndTime Setter
// 辅营创建结束时间
func (r *AlitripTripvpAgentOrderSearchRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r AlitripTripvpAgentOrderSearchRequest) GetEndTime() string {
    return r.endTime
}
// OrderStatus Setter
// 订单状态，1-待支付 2-支付成功 3-	辅营出货成功 4-订单取消
func (r *AlitripTripvpAgentOrderSearchRequest) SetOrderStatus(orderStatus int64) error {
    r.orderStatus = orderStatus
    r.Set("order_status", orderStatus)
    return nil
}

// OrderStatus Getter
func (r AlitripTripvpAgentOrderSearchRequest) GetOrderStatus() int64 {
    return r.orderStatus
}
// PageSize Setter
// 分页行数
func (r *AlitripTripvpAgentOrderSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlitripTripvpAgentOrderSearchRequest) GetPageSize() int64 {
    return r.pageSize
}
