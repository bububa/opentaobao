package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家查询退票申请 API请求
taobao.alitrip.ie.agent.refund.search

卖家查询退票申请
*/
type TaobaoAlitripIeAgentRefundSearchRequest struct {
    model.Params
    // 查询起始时间
    _createStartTime   string
    // 查询结束时间
    _createEndTime   string
    // WAIT(1,"待处理"), AGREED(2, "已同意"),REFUSE(3, "已拒绝"),PROCESS(6, "已受理"), SUCCESS(7, "已退款");
    _refundStatus   int64
    // 从1开始
    _pageIndex   int64
    // 每页大小
    _pageSize   int64
    // 代理商id
    _agentId   int64
}

// 初始化TaobaoAlitripIeAgentRefundSearchRequest对象
func NewTaobaoAlitripIeAgentRefundSearchRequest() *TaobaoAlitripIeAgentRefundSearchRequest{
    return &TaobaoAlitripIeAgentRefundSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundSearchRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreateStartTime Setter
// 查询起始时间
func (r *TaobaoAlitripIeAgentRefundSearchRequest) SetCreateStartTime(_createStartTime string) error {
    r._createStartTime = _createStartTime
    r.Set("create_start_time", _createStartTime)
    return nil
}

// CreateStartTime Getter
func (r TaobaoAlitripIeAgentRefundSearchRequest) GetCreateStartTime() string {
    return r._createStartTime
}
// CreateEndTime Setter
// 查询结束时间
func (r *TaobaoAlitripIeAgentRefundSearchRequest) SetCreateEndTime(_createEndTime string) error {
    r._createEndTime = _createEndTime
    r.Set("create_end_time", _createEndTime)
    return nil
}

// CreateEndTime Getter
func (r TaobaoAlitripIeAgentRefundSearchRequest) GetCreateEndTime() string {
    return r._createEndTime
}
// RefundStatus Setter
// WAIT(1,"待处理"), AGREED(2, "已同意"),REFUSE(3, "已拒绝"),PROCESS(6, "已受理"), SUCCESS(7, "已退款");
func (r *TaobaoAlitripIeAgentRefundSearchRequest) SetRefundStatus(_refundStatus int64) error {
    r._refundStatus = _refundStatus
    r.Set("refund_status", _refundStatus)
    return nil
}

// RefundStatus Getter
func (r TaobaoAlitripIeAgentRefundSearchRequest) GetRefundStatus() int64 {
    return r._refundStatus
}
// PageIndex Setter
// 从1开始
func (r *TaobaoAlitripIeAgentRefundSearchRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoAlitripIeAgentRefundSearchRequest) GetPageIndex() int64 {
    return r._pageIndex
}
// PageSize Setter
// 每页大小
func (r *TaobaoAlitripIeAgentRefundSearchRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoAlitripIeAgentRefundSearchRequest) GetPageSize() int64 {
    return r._pageSize
}
// AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundSearchRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoAlitripIeAgentRefundSearchRequest) GetAgentId() int64 {
    return r._agentId
}
