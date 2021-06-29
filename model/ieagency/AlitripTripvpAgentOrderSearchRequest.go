package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】查询辅营订单列表 APIRequest
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

func NewAlitripTripvpAgentOrderSearchRequest() *AlitripTripvpAgentOrderSearchRequest{
    return &AlitripTripvpAgentOrderSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTripvpAgentOrderSearchRequest) GetApiMethodName() string {
    return "alitrip.tripvp.agent.order.search"
}

func (r AlitripTripvpAgentOrderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTripvpAgentOrderSearchRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r AlitripTripvpAgentOrderSearchRequest) GetAgentId() int64 {
    return r.agentId
}

func (r *AlitripTripvpAgentOrderSearchRequest) SetBeginTime(beginTime string) error {
    r.beginTime = beginTime
    r.Set("begin_time", beginTime)
    return nil
}

func (r AlitripTripvpAgentOrderSearchRequest) GetBeginTime() string {
    return r.beginTime
}

func (r *AlitripTripvpAgentOrderSearchRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlitripTripvpAgentOrderSearchRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *AlitripTripvpAgentOrderSearchRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r AlitripTripvpAgentOrderSearchRequest) GetEndTime() string {
    return r.endTime
}

func (r *AlitripTripvpAgentOrderSearchRequest) SetOrderStatus(orderStatus int64) error {
    r.orderStatus = orderStatus
    r.Set("order_status", orderStatus)
    return nil
}

func (r AlitripTripvpAgentOrderSearchRequest) GetOrderStatus() int64 {
    return r.orderStatus
}

func (r *AlitripTripvpAgentOrderSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlitripTripvpAgentOrderSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

