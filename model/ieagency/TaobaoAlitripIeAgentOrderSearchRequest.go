package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】订单列表查询 API请求
taobao.alitrip.ie.agent.order.search

根据指定条件查询国际机票订单列表
*/
type TaobaoAlitripIeAgentOrderSearchRequest struct {
    model.Params
    // 代理商ID
    _agentId   int64
    // 订单起始日期
    _beginTime   string
    // 当前页码
    _currentPage   int64
    // 订单结束日期
    _endTime   string
    // 订单状态（只能传入一个状态，不支持多个一起传）
    _orderStatus   string
    // 分页大小
    _pageSize   int64
    // 0:自有运价;3:公布运价;9:大卖家API;11私有库存
    _fareSource   int64
    // 供应渠道/资源码
    _resourceCode   string
    // officeNo
    _officeNo   string
}

// 初始化TaobaoAlitripIeAgentOrderSearchRequest对象
func NewTaobaoAlitripIeAgentOrderSearchRequest() *TaobaoAlitripIeAgentOrderSearchRequest{
    return &TaobaoAlitripIeAgentOrderSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentOrderSearchRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.order.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentOrderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 代理商ID
func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoAlitripIeAgentOrderSearchRequest) GetAgentId() int64 {
    return r._agentId
}
// BeginTime Setter
// 订单起始日期
func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetBeginTime(_beginTime string) error {
    r._beginTime = _beginTime
    r.Set("begin_time", _beginTime)
    return nil
}

// BeginTime Getter
func (r TaobaoAlitripIeAgentOrderSearchRequest) GetBeginTime() string {
    return r._beginTime
}
// CurrentPage Setter
// 当前页码
func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoAlitripIeAgentOrderSearchRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// EndTime Setter
// 订单结束日期
func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoAlitripIeAgentOrderSearchRequest) GetEndTime() string {
    return r._endTime
}
// OrderStatus Setter
// 订单状态（只能传入一个状态，不支持多个一起传）
func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetOrderStatus(_orderStatus string) error {
    r._orderStatus = _orderStatus
    r.Set("order_status", _orderStatus)
    return nil
}

// OrderStatus Getter
func (r TaobaoAlitripIeAgentOrderSearchRequest) GetOrderStatus() string {
    return r._orderStatus
}
// PageSize Setter
// 分页大小
func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoAlitripIeAgentOrderSearchRequest) GetPageSize() int64 {
    return r._pageSize
}
// FareSource Setter
// 0:自有运价;3:公布运价;9:大卖家API;11私有库存
func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetFareSource(_fareSource int64) error {
    r._fareSource = _fareSource
    r.Set("fare_source", _fareSource)
    return nil
}

// FareSource Getter
func (r TaobaoAlitripIeAgentOrderSearchRequest) GetFareSource() int64 {
    return r._fareSource
}
// ResourceCode Setter
// 供应渠道/资源码
func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetResourceCode(_resourceCode string) error {
    r._resourceCode = _resourceCode
    r.Set("resource_code", _resourceCode)
    return nil
}

// ResourceCode Getter
func (r TaobaoAlitripIeAgentOrderSearchRequest) GetResourceCode() string {
    return r._resourceCode
}
// OfficeNo Setter
// officeNo
func (r *TaobaoAlitripIeAgentOrderSearchRequest) SetOfficeNo(_officeNo string) error {
    r._officeNo = _officeNo
    r.Set("office_no", _officeNo)
    return nil
}

// OfficeNo Getter
func (r TaobaoAlitripIeAgentOrderSearchRequest) GetOfficeNo() string {
    return r._officeNo
}
