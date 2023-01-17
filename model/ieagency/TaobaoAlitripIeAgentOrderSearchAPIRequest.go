package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentOrderSearchAPIRequest 【国际机票】订单列表查询 API请求
// taobao.alitrip.ie.agent.order.search
//
// 根据指定条件查询国际机票订单列表
type TaobaoAlitripIeAgentOrderSearchAPIRequest struct {
	model.Params
	// 订单起始日期
	_beginTime string
	// 订单结束日期
	_endTime string
	// 订单状态（只能传入一个状态，不支持多个一起传）
	_orderStatus string
	// 供应渠道/资源码
	_resourceCode string
	// officeNo
	_officeNo string
	// 代理商ID
	_agentId int64
	// 当前页码
	_currentPage int64
	// 分页大小
	_pageSize int64
	// 0:自有运价;3:公布运价;9:大卖家API;11私有库存
	_fareSource int64
}

// NewTaobaoAlitripIeAgentOrderSearchRequest 初始化TaobaoAlitripIeAgentOrderSearchAPIRequest对象
func NewTaobaoAlitripIeAgentOrderSearchRequest() *TaobaoAlitripIeAgentOrderSearchAPIRequest {
	return &TaobaoAlitripIeAgentOrderSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentOrderSearchAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentOrderSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripIeAgentOrderSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBeginTime is BeginTime Setter
// 订单起始日期
func (r *TaobaoAlitripIeAgentOrderSearchAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// GetBeginTime BeginTime Getter
func (r TaobaoAlitripIeAgentOrderSearchAPIRequest) GetBeginTime() string {
	return r._beginTime
}

// SetEndTime is EndTime Setter
// 订单结束日期
func (r *TaobaoAlitripIeAgentOrderSearchAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoAlitripIeAgentOrderSearchAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetOrderStatus is OrderStatus Setter
// 订单状态（只能传入一个状态，不支持多个一起传）
func (r *TaobaoAlitripIeAgentOrderSearchAPIRequest) SetOrderStatus(_orderStatus string) error {
	r._orderStatus = _orderStatus
	r.Set("order_status", _orderStatus)
	return nil
}

// GetOrderStatus OrderStatus Getter
func (r TaobaoAlitripIeAgentOrderSearchAPIRequest) GetOrderStatus() string {
	return r._orderStatus
}

// SetResourceCode is ResourceCode Setter
// 供应渠道/资源码
func (r *TaobaoAlitripIeAgentOrderSearchAPIRequest) SetResourceCode(_resourceCode string) error {
	r._resourceCode = _resourceCode
	r.Set("resource_code", _resourceCode)
	return nil
}

// GetResourceCode ResourceCode Getter
func (r TaobaoAlitripIeAgentOrderSearchAPIRequest) GetResourceCode() string {
	return r._resourceCode
}

// SetOfficeNo is OfficeNo Setter
// officeNo
func (r *TaobaoAlitripIeAgentOrderSearchAPIRequest) SetOfficeNo(_officeNo string) error {
	r._officeNo = _officeNo
	r.Set("office_no", _officeNo)
	return nil
}

// GetOfficeNo OfficeNo Getter
func (r TaobaoAlitripIeAgentOrderSearchAPIRequest) GetOfficeNo() string {
	return r._officeNo
}

// SetAgentId is AgentId Setter
// 代理商ID
func (r *TaobaoAlitripIeAgentOrderSearchAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoAlitripIeAgentOrderSearchAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetCurrentPage is CurrentPage Setter
// 当前页码
func (r *TaobaoAlitripIeAgentOrderSearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoAlitripIeAgentOrderSearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TaobaoAlitripIeAgentOrderSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoAlitripIeAgentOrderSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetFareSource is FareSource Setter
// 0:自有运价;3:公布运价;9:大卖家API;11私有库存
func (r *TaobaoAlitripIeAgentOrderSearchAPIRequest) SetFareSource(_fareSource int64) error {
	r._fareSource = _fareSource
	r.Set("fare_source", _fareSource)
	return nil
}

// GetFareSource FareSource Getter
func (r TaobaoAlitripIeAgentOrderSearchAPIRequest) GetFareSource() int64 {
	return r._fareSource
}
