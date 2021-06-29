package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】订单列表查询 APIRequest
taobao.alitrip.ie.agent.order.search

根据指定条件查询国际机票订单列表
*/
type TaobaoAlitripIeAgentOrderSearchRequest struct {
    model.Params

    // 代理商ID
    agentId   int64 

    // 订单起始日期
    beginTime   string 

    // 当前页码
    currentPage   int64 

    // 订单结束日期
    endTime   string 

    // 订单状态（只能传入一个状态，不支持多个一起传）
    orderStatus   string 

    // 分页大小
    pageSize   int64 

    // 0:自有运价;3:公布运价;9:大卖家API;11私有库存
    fareSource   int64 

    // 供应渠道/资源码
    resourceCode   string 

    // officeNo
    officeNo   string 

}

func NewTaobaoAlitripIeAgentOrderSearchRequest() *TaobaoAlitripIeAgentOrderSearchRequest{
    return &TaobaoAlitripIeAgentOrderSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripIeAgentOrderSearchRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.order.search"
}

func (r TaobaoAlitripIeAgentOrderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoAlitripIeAgentOrderSearchRequest) GetAgentId() int64 {
    return r.agentId
}

func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetBeginTime(beginTime string) error {
    r.beginTime = beginTime
    r.Set("begin_time", beginTime)
    return nil
}

func (r TaobaoAlitripIeAgentOrderSearchRequest) GetBeginTime() string {
    return r.beginTime
}

func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r TaobaoAlitripIeAgentOrderSearchRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoAlitripIeAgentOrderSearchRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetOrderStatus(orderStatus string) error {
    r.orderStatus = orderStatus
    r.Set("order_status", orderStatus)
    return nil
}

func (r TaobaoAlitripIeAgentOrderSearchRequest) GetOrderStatus() string {
    return r.orderStatus
}

func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoAlitripIeAgentOrderSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetFareSource(fareSource int64) error {
    r.fareSource = fareSource
    r.Set("fare_source", fareSource)
    return nil
}

func (r TaobaoAlitripIeAgentOrderSearchRequest) GetFareSource() int64 {
    return r.fareSource
}

func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetResourceCode(resourceCode string) error {
    r.resourceCode = resourceCode
    r.Set("resource_code", resourceCode)
    return nil
}

func (r TaobaoAlitripIeAgentOrderSearchRequest) GetResourceCode() string {
    return r.resourceCode
}

func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetOfficeNo(officeNo string) error {
    r.officeNo = officeNo
    r.Set("office_no", officeNo)
    return nil
}

func (r TaobaoAlitripIeAgentOrderSearchRequest) GetOfficeNo() string {
    return r.officeNo
}

