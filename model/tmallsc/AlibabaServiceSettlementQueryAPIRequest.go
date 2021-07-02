package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServiceSettlementQueryAPIRequest 服务平台结算单明细查询服务 API请求
// alibaba.service.settlement.query
//
// 给服务商提供结算单明细查询功能
type AlibabaServiceSettlementQueryAPIRequest struct {
	model.Params
	// 账单查询开始时间。格式示例 2019-03-26 17:15:28
	_gmtCreateStart string
	// 账单查询结束时间，时间区间限制未15分钟。 格式示例 2019-03-26 17:15:28
	_gmtCreateEnd string
	// 当前页面，开始值为1
	_currentPage int64
	// 页面展示条数大小
	_pageSize int64
	// 工单ID
	_workcardId int64
	// 交易主订单号码
	_parentTradeOrderId int64
	// 服务单号
	_serviceOrderId int64
	// 交易实物订单号
	_masterTradeOrderId int64
	// 交易服务订单号
	_serviceTradeOrderId int64
	// 账单修改开始时间。
	_gmtModifiedEnd string
	// 账单修改结束时间，时间区间限制未15分钟。
	_gmtModifiedStart string
}

// NewAlibabaServiceSettlementQueryRequest 初始化AlibabaServiceSettlementQueryAPIRequest对象
func NewAlibabaServiceSettlementQueryRequest() *AlibabaServiceSettlementQueryAPIRequest {
	return &AlibabaServiceSettlementQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServiceSettlementQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.service.settlement.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServiceSettlementQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetGmtCreateStart is GmtCreateStart Setter
// 账单查询开始时间。格式示例 2019-03-26 17:15:28
func (r *AlibabaServiceSettlementQueryAPIRequest) SetGmtCreateStart(_gmtCreateStart string) error {
	r._gmtCreateStart = _gmtCreateStart
	r.Set("gmt_create_start", _gmtCreateStart)
	return nil
}

// GetGmtCreateStart GmtCreateStart Getter
func (r AlibabaServiceSettlementQueryAPIRequest) GetGmtCreateStart() string {
	return r._gmtCreateStart
}

// SetGmtCreateEnd is GmtCreateEnd Setter
// 账单查询结束时间，时间区间限制未15分钟。 格式示例 2019-03-26 17:15:28
func (r *AlibabaServiceSettlementQueryAPIRequest) SetGmtCreateEnd(_gmtCreateEnd string) error {
	r._gmtCreateEnd = _gmtCreateEnd
	r.Set("gmt_create_end", _gmtCreateEnd)
	return nil
}

// GetGmtCreateEnd GmtCreateEnd Getter
func (r AlibabaServiceSettlementQueryAPIRequest) GetGmtCreateEnd() string {
	return r._gmtCreateEnd
}

// SetCurrentPage is CurrentPage Setter
// 当前页面，开始值为1
func (r *AlibabaServiceSettlementQueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaServiceSettlementQueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 页面展示条数大小
func (r *AlibabaServiceSettlementQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaServiceSettlementQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetWorkcardId is WorkcardId Setter
// 工单ID
func (r *AlibabaServiceSettlementQueryAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r AlibabaServiceSettlementQueryAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetParentTradeOrderId is ParentTradeOrderId Setter
// 交易主订单号码
func (r *AlibabaServiceSettlementQueryAPIRequest) SetParentTradeOrderId(_parentTradeOrderId int64) error {
	r._parentTradeOrderId = _parentTradeOrderId
	r.Set("parent_trade_order_id", _parentTradeOrderId)
	return nil
}

// GetParentTradeOrderId ParentTradeOrderId Getter
func (r AlibabaServiceSettlementQueryAPIRequest) GetParentTradeOrderId() int64 {
	return r._parentTradeOrderId
}

// SetServiceOrderId is ServiceOrderId Setter
// 服务单号
func (r *AlibabaServiceSettlementQueryAPIRequest) SetServiceOrderId(_serviceOrderId int64) error {
	r._serviceOrderId = _serviceOrderId
	r.Set("service_order_id", _serviceOrderId)
	return nil
}

// GetServiceOrderId ServiceOrderId Getter
func (r AlibabaServiceSettlementQueryAPIRequest) GetServiceOrderId() int64 {
	return r._serviceOrderId
}

// SetMasterTradeOrderId is MasterTradeOrderId Setter
// 交易实物订单号
func (r *AlibabaServiceSettlementQueryAPIRequest) SetMasterTradeOrderId(_masterTradeOrderId int64) error {
	r._masterTradeOrderId = _masterTradeOrderId
	r.Set("master_trade_order_id", _masterTradeOrderId)
	return nil
}

// GetMasterTradeOrderId MasterTradeOrderId Getter
func (r AlibabaServiceSettlementQueryAPIRequest) GetMasterTradeOrderId() int64 {
	return r._masterTradeOrderId
}

// SetServiceTradeOrderId is ServiceTradeOrderId Setter
// 交易服务订单号
func (r *AlibabaServiceSettlementQueryAPIRequest) SetServiceTradeOrderId(_serviceTradeOrderId int64) error {
	r._serviceTradeOrderId = _serviceTradeOrderId
	r.Set("service_trade_order_id", _serviceTradeOrderId)
	return nil
}

// GetServiceTradeOrderId ServiceTradeOrderId Getter
func (r AlibabaServiceSettlementQueryAPIRequest) GetServiceTradeOrderId() int64 {
	return r._serviceTradeOrderId
}

// SetGmtModifiedEnd is GmtModifiedEnd Setter
// 账单修改开始时间。
func (r *AlibabaServiceSettlementQueryAPIRequest) SetGmtModifiedEnd(_gmtModifiedEnd string) error {
	r._gmtModifiedEnd = _gmtModifiedEnd
	r.Set("gmt_modified_end", _gmtModifiedEnd)
	return nil
}

// GetGmtModifiedEnd GmtModifiedEnd Getter
func (r AlibabaServiceSettlementQueryAPIRequest) GetGmtModifiedEnd() string {
	return r._gmtModifiedEnd
}

// SetGmtModifiedStart is GmtModifiedStart Setter
// 账单修改结束时间，时间区间限制未15分钟。
func (r *AlibabaServiceSettlementQueryAPIRequest) SetGmtModifiedStart(_gmtModifiedStart string) error {
	r._gmtModifiedStart = _gmtModifiedStart
	r.Set("gmt_modified_start", _gmtModifiedStart)
	return nil
}

// GetGmtModifiedStart GmtModifiedStart Getter
func (r AlibabaServiceSettlementQueryAPIRequest) GetGmtModifiedStart() string {
	return r._gmtModifiedStart
}
