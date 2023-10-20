package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicesettlementqueryAPIRequest 服务平台结算单明细查询服务 API请求
// alibaba.service.settlement.query
//
// 给服务商提供结算单明细查询功能
type AlibabaservicesettlementqueryAPIRequest struct {
	model.Params
	// 账单查询开始时间。格式示例 2019-03-26 17:15:28
	_gmtCreateStart string
	// 账单查询结束时间，时间区间限制未15分钟。 格式示例 2019-03-26 17:15:28
	_gmtCreateEnd string
	// 账单修改开始时间。
	_gmtModifiedEnd string
	// 账单修改结束时间，时间区间限制未15分钟。
	_gmtModifiedStart string
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
}

// NewAlibabaservicesettlementqueryRequest 初始化AlibabaservicesettlementqueryAPIRequest对象
func NewAlibabaservicesettlementqueryRequest() *AlibabaservicesettlementqueryAPIRequest {
	return &AlibabaservicesettlementqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaservicesettlementqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.service.settlement.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaservicesettlementqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaservicesettlementqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGmtCreateStart is GmtCreateStart Setter
// 账单查询开始时间。格式示例 2019-03-26 17:15:28
func (r *AlibabaservicesettlementqueryAPIRequest) SetGmtCreateStart(_gmtCreateStart string) error {
	r._gmtCreateStart = _gmtCreateStart
	r.Set("gmt_create_start", _gmtCreateStart)
	return nil
}

// GetGmtCreateStart GmtCreateStart Getter
func (r AlibabaservicesettlementqueryAPIRequest) GetGmtCreateStart() string {
	return r._gmtCreateStart
}

// SetGmtCreateEnd is GmtCreateEnd Setter
// 账单查询结束时间，时间区间限制未15分钟。 格式示例 2019-03-26 17:15:28
func (r *AlibabaservicesettlementqueryAPIRequest) SetGmtCreateEnd(_gmtCreateEnd string) error {
	r._gmtCreateEnd = _gmtCreateEnd
	r.Set("gmt_create_end", _gmtCreateEnd)
	return nil
}

// GetGmtCreateEnd GmtCreateEnd Getter
func (r AlibabaservicesettlementqueryAPIRequest) GetGmtCreateEnd() string {
	return r._gmtCreateEnd
}

// SetGmtModifiedEnd is GmtModifiedEnd Setter
// 账单修改开始时间。
func (r *AlibabaservicesettlementqueryAPIRequest) SetGmtModifiedEnd(_gmtModifiedEnd string) error {
	r._gmtModifiedEnd = _gmtModifiedEnd
	r.Set("gmt_modified_end", _gmtModifiedEnd)
	return nil
}

// GetGmtModifiedEnd GmtModifiedEnd Getter
func (r AlibabaservicesettlementqueryAPIRequest) GetGmtModifiedEnd() string {
	return r._gmtModifiedEnd
}

// SetGmtModifiedStart is GmtModifiedStart Setter
// 账单修改结束时间，时间区间限制未15分钟。
func (r *AlibabaservicesettlementqueryAPIRequest) SetGmtModifiedStart(_gmtModifiedStart string) error {
	r._gmtModifiedStart = _gmtModifiedStart
	r.Set("gmt_modified_start", _gmtModifiedStart)
	return nil
}

// GetGmtModifiedStart GmtModifiedStart Getter
func (r AlibabaservicesettlementqueryAPIRequest) GetGmtModifiedStart() string {
	return r._gmtModifiedStart
}

// SetCurrentPage is CurrentPage Setter
// 当前页面，开始值为1
func (r *AlibabaservicesettlementqueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaservicesettlementqueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 页面展示条数大小
func (r *AlibabaservicesettlementqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaservicesettlementqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetWorkcardId is WorkcardId Setter
// 工单ID
func (r *AlibabaservicesettlementqueryAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r AlibabaservicesettlementqueryAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetParentTradeOrderId is ParentTradeOrderId Setter
// 交易主订单号码
func (r *AlibabaservicesettlementqueryAPIRequest) SetParentTradeOrderId(_parentTradeOrderId int64) error {
	r._parentTradeOrderId = _parentTradeOrderId
	r.Set("parent_trade_order_id", _parentTradeOrderId)
	return nil
}

// GetParentTradeOrderId ParentTradeOrderId Getter
func (r AlibabaservicesettlementqueryAPIRequest) GetParentTradeOrderId() int64 {
	return r._parentTradeOrderId
}

// SetServiceOrderId is ServiceOrderId Setter
// 服务单号
func (r *AlibabaservicesettlementqueryAPIRequest) SetServiceOrderId(_serviceOrderId int64) error {
	r._serviceOrderId = _serviceOrderId
	r.Set("service_order_id", _serviceOrderId)
	return nil
}

// GetServiceOrderId ServiceOrderId Getter
func (r AlibabaservicesettlementqueryAPIRequest) GetServiceOrderId() int64 {
	return r._serviceOrderId
}

// SetMasterTradeOrderId is MasterTradeOrderId Setter
// 交易实物订单号
func (r *AlibabaservicesettlementqueryAPIRequest) SetMasterTradeOrderId(_masterTradeOrderId int64) error {
	r._masterTradeOrderId = _masterTradeOrderId
	r.Set("master_trade_order_id", _masterTradeOrderId)
	return nil
}

// GetMasterTradeOrderId MasterTradeOrderId Getter
func (r AlibabaservicesettlementqueryAPIRequest) GetMasterTradeOrderId() int64 {
	return r._masterTradeOrderId
}

// SetServiceTradeOrderId is ServiceTradeOrderId Setter
// 交易服务订单号
func (r *AlibabaservicesettlementqueryAPIRequest) SetServiceTradeOrderId(_serviceTradeOrderId int64) error {
	r._serviceTradeOrderId = _serviceTradeOrderId
	r.Set("service_trade_order_id", _serviceTradeOrderId)
	return nil
}

// GetServiceTradeOrderId ServiceTradeOrderId Getter
func (r AlibabaservicesettlementqueryAPIRequest) GetServiceTradeOrderId() int64 {
	return r._serviceTradeOrderId
}
