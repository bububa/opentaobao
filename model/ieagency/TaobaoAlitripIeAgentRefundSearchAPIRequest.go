package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripieagentrefundsearchAPIRequest 卖家查询退票申请 API请求
// taobao.alitrip.ie.agent.refund.search
//
// 卖家查询退票申请
type TaobaoalitripieagentrefundsearchAPIRequest struct {
	model.Params
	// 查询起始时间
	_createStartTime string
	// 查询结束时间
	_createEndTime string
	// WAIT(1,"待处理"), AGREED(2, "已同意"),REFUSE(3, "已拒绝"),PROCESS(6, "已受理"), SUCCESS(7, "已退款");
	_refundStatus int64
	// 从1开始
	_pageIndex int64
	// 每页大小
	_pageSize int64
	// 代理商id
	_agentId int64
}

// NewTaobaoalitripieagentrefundsearchRequest 初始化TaobaoalitripieagentrefundsearchAPIRequest对象
func NewTaobaoalitripieagentrefundsearchRequest() *TaobaoalitripieagentrefundsearchAPIRequest {
	return &TaobaoalitripieagentrefundsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripieagentrefundsearchAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripieagentrefundsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripieagentrefundsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateStartTime is CreateStartTime Setter
// 查询起始时间
func (r *TaobaoalitripieagentrefundsearchAPIRequest) SetCreateStartTime(_createStartTime string) error {
	r._createStartTime = _createStartTime
	r.Set("create_start_time", _createStartTime)
	return nil
}

// GetCreateStartTime CreateStartTime Getter
func (r TaobaoalitripieagentrefundsearchAPIRequest) GetCreateStartTime() string {
	return r._createStartTime
}

// SetCreateEndTime is CreateEndTime Setter
// 查询结束时间
func (r *TaobaoalitripieagentrefundsearchAPIRequest) SetCreateEndTime(_createEndTime string) error {
	r._createEndTime = _createEndTime
	r.Set("create_end_time", _createEndTime)
	return nil
}

// GetCreateEndTime CreateEndTime Getter
func (r TaobaoalitripieagentrefundsearchAPIRequest) GetCreateEndTime() string {
	return r._createEndTime
}

// SetRefundStatus is RefundStatus Setter
// WAIT(1,&#34;待处理&#34;), AGREED(2, &#34;已同意&#34;),REFUSE(3, &#34;已拒绝&#34;),PROCESS(6, &#34;已受理&#34;), SUCCESS(7, &#34;已退款&#34;);
func (r *TaobaoalitripieagentrefundsearchAPIRequest) SetRefundStatus(_refundStatus int64) error {
	r._refundStatus = _refundStatus
	r.Set("refund_status", _refundStatus)
	return nil
}

// GetRefundStatus RefundStatus Getter
func (r TaobaoalitripieagentrefundsearchAPIRequest) GetRefundStatus() int64 {
	return r._refundStatus
}

// SetPageIndex is PageIndex Setter
// 从1开始
func (r *TaobaoalitripieagentrefundsearchAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoalitripieagentrefundsearchAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaoalitripieagentrefundsearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoalitripieagentrefundsearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoalitripieagentrefundsearchAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoalitripieagentrefundsearchAPIRequest) GetAgentId() int64 {
	return r._agentId
}
