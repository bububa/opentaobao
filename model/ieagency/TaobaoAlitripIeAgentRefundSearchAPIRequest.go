package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundSearchAPIRequest 卖家查询退票申请 API请求
// taobao.alitrip.ie.agent.refund.search
//
// 卖家查询退票申请
type TaobaoAlitripIeAgentRefundSearchAPIRequest struct {
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

// NewTaobaoAlitripIeAgentRefundSearchRequest 初始化TaobaoAlitripIeAgentRefundSearchAPIRequest对象
func NewTaobaoAlitripIeAgentRefundSearchRequest() *TaobaoAlitripIeAgentRefundSearchAPIRequest {
	return &TaobaoAlitripIeAgentRefundSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundSearchAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripIeAgentRefundSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateStartTime is CreateStartTime Setter
// 查询起始时间
func (r *TaobaoAlitripIeAgentRefundSearchAPIRequest) SetCreateStartTime(_createStartTime string) error {
	r._createStartTime = _createStartTime
	r.Set("create_start_time", _createStartTime)
	return nil
}

// GetCreateStartTime CreateStartTime Getter
func (r TaobaoAlitripIeAgentRefundSearchAPIRequest) GetCreateStartTime() string {
	return r._createStartTime
}

// SetCreateEndTime is CreateEndTime Setter
// 查询结束时间
func (r *TaobaoAlitripIeAgentRefundSearchAPIRequest) SetCreateEndTime(_createEndTime string) error {
	r._createEndTime = _createEndTime
	r.Set("create_end_time", _createEndTime)
	return nil
}

// GetCreateEndTime CreateEndTime Getter
func (r TaobaoAlitripIeAgentRefundSearchAPIRequest) GetCreateEndTime() string {
	return r._createEndTime
}

// SetRefundStatus is RefundStatus Setter
// WAIT(1,&#34;待处理&#34;), AGREED(2, &#34;已同意&#34;),REFUSE(3, &#34;已拒绝&#34;),PROCESS(6, &#34;已受理&#34;), SUCCESS(7, &#34;已退款&#34;);
func (r *TaobaoAlitripIeAgentRefundSearchAPIRequest) SetRefundStatus(_refundStatus int64) error {
	r._refundStatus = _refundStatus
	r.Set("refund_status", _refundStatus)
	return nil
}

// GetRefundStatus RefundStatus Getter
func (r TaobaoAlitripIeAgentRefundSearchAPIRequest) GetRefundStatus() int64 {
	return r._refundStatus
}

// SetPageIndex is PageIndex Setter
// 从1开始
func (r *TaobaoAlitripIeAgentRefundSearchAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoAlitripIeAgentRefundSearchAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaoAlitripIeAgentRefundSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoAlitripIeAgentRefundSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundSearchAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoAlitripIeAgentRefundSearchAPIRequest) GetAgentId() int64 {
	return r._agentId
}
