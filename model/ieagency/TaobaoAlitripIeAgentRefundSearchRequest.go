package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家查询退票申请 APIRequest
taobao.alitrip.ie.agent.refund.search

卖家查询退票申请
*/
type TaobaoAlitripIeAgentRefundSearchRequest struct {
    model.Params

    // 查询起始时间
    createStartTime   string 

    // 查询结束时间
    createEndTime   string 

    // WAIT(1,"待处理"), AGREED(2, "已同意"),REFUSE(3, "已拒绝"),PROCESS(6, "已受理"), SUCCESS(7, "已退款");
    refundStatus   int64 

    // 从1开始
    pageIndex   int64 

    // 每页大小
    pageSize   int64 

    // 代理商id
    agentId   int64 

}

func NewTaobaoAlitripIeAgentRefundSearchRequest() *TaobaoAlitripIeAgentRefundSearchRequest{
    return &TaobaoAlitripIeAgentRefundSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripIeAgentRefundSearchRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.search"
}

func (r TaobaoAlitripIeAgentRefundSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripIeAgentRefundSearchRequest) SetCreateStartTime(createStartTime string) error {
    r.createStartTime = createStartTime
    r.Set("create_start_time", createStartTime)
    return nil
}

func (r TaobaoAlitripIeAgentRefundSearchRequest) GetCreateStartTime() string {
    return r.createStartTime
}

func (r *TaobaoAlitripIeAgentRefundSearchRequest) SetCreateEndTime(createEndTime string) error {
    r.createEndTime = createEndTime
    r.Set("create_end_time", createEndTime)
    return nil
}

func (r TaobaoAlitripIeAgentRefundSearchRequest) GetCreateEndTime() string {
    return r.createEndTime
}

func (r *TaobaoAlitripIeAgentRefundSearchRequest) SetRefundStatus(refundStatus int64) error {
    r.refundStatus = refundStatus
    r.Set("refund_status", refundStatus)
    return nil
}

func (r TaobaoAlitripIeAgentRefundSearchRequest) GetRefundStatus() int64 {
    return r.refundStatus
}

func (r *TaobaoAlitripIeAgentRefundSearchRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r TaobaoAlitripIeAgentRefundSearchRequest) GetPageIndex() int64 {
    return r.pageIndex
}

func (r *TaobaoAlitripIeAgentRefundSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoAlitripIeAgentRefundSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoAlitripIeAgentRefundSearchRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoAlitripIeAgentRefundSearchRequest) GetAgentId() int64 {
    return r.agentId
}

