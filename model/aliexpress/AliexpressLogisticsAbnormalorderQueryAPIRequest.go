package aliexpress

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresslogisticsabnormalorderqueryAPIRequest 异常订单查询 API请求
// aliexpress.logistics.abnormalorder.query
//
// 异常订单查询
type AliexpresslogisticsabnormalorderqueryAPIRequest struct {
	model.Params
	// 订单创建时间开始
	_gmtCreateStart string
	// 状态编码列表
	_warehouseStatusList string
	// 状态变更时间开始
	_gmtStatusUpdateStart string
	// 运单号
	_intlTrackingNo string
	// 订单创建时间结束
	_gmtCreateEnd string
	// 状态变更时间结束
	_gmtStatusUpdateEnd string
	// 交易订单号
	_tradeOrderId int64
	// 页大小
	_pageSize int64
	// 当前页
	_currentPage int64
}

// NewAliexpresslogisticsabnormalorderqueryRequest 初始化AliexpresslogisticsabnormalorderqueryAPIRequest对象
func NewAliexpresslogisticsabnormalorderqueryRequest() *AliexpresslogisticsabnormalorderqueryAPIRequest {
	return &AliexpresslogisticsabnormalorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresslogisticsabnormalorderqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.logistics.abnormalorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresslogisticsabnormalorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresslogisticsabnormalorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGmtCreateStart is GmtCreateStart Setter
// 订单创建时间开始
func (r *AliexpresslogisticsabnormalorderqueryAPIRequest) SetGmtCreateStart(_gmtCreateStart string) error {
	r._gmtCreateStart = _gmtCreateStart
	r.Set("gmt_create_start", _gmtCreateStart)
	return nil
}

// GetGmtCreateStart GmtCreateStart Getter
func (r AliexpresslogisticsabnormalorderqueryAPIRequest) GetGmtCreateStart() string {
	return r._gmtCreateStart
}

// SetWarehouseStatusList is WarehouseStatusList Setter
// 状态编码列表
func (r *AliexpresslogisticsabnormalorderqueryAPIRequest) SetWarehouseStatusList(_warehouseStatusList string) error {
	r._warehouseStatusList = _warehouseStatusList
	r.Set("warehouse_status_list", _warehouseStatusList)
	return nil
}

// GetWarehouseStatusList WarehouseStatusList Getter
func (r AliexpresslogisticsabnormalorderqueryAPIRequest) GetWarehouseStatusList() string {
	return r._warehouseStatusList
}

// SetGmtStatusUpdateStart is GmtStatusUpdateStart Setter
// 状态变更时间开始
func (r *AliexpresslogisticsabnormalorderqueryAPIRequest) SetGmtStatusUpdateStart(_gmtStatusUpdateStart string) error {
	r._gmtStatusUpdateStart = _gmtStatusUpdateStart
	r.Set("gmt_status_update_start", _gmtStatusUpdateStart)
	return nil
}

// GetGmtStatusUpdateStart GmtStatusUpdateStart Getter
func (r AliexpresslogisticsabnormalorderqueryAPIRequest) GetGmtStatusUpdateStart() string {
	return r._gmtStatusUpdateStart
}

// SetIntlTrackingNo is IntlTrackingNo Setter
// 运单号
func (r *AliexpresslogisticsabnormalorderqueryAPIRequest) SetIntlTrackingNo(_intlTrackingNo string) error {
	r._intlTrackingNo = _intlTrackingNo
	r.Set("intl_tracking_no", _intlTrackingNo)
	return nil
}

// GetIntlTrackingNo IntlTrackingNo Getter
func (r AliexpresslogisticsabnormalorderqueryAPIRequest) GetIntlTrackingNo() string {
	return r._intlTrackingNo
}

// SetGmtCreateEnd is GmtCreateEnd Setter
// 订单创建时间结束
func (r *AliexpresslogisticsabnormalorderqueryAPIRequest) SetGmtCreateEnd(_gmtCreateEnd string) error {
	r._gmtCreateEnd = _gmtCreateEnd
	r.Set("gmt_create_end", _gmtCreateEnd)
	return nil
}

// GetGmtCreateEnd GmtCreateEnd Getter
func (r AliexpresslogisticsabnormalorderqueryAPIRequest) GetGmtCreateEnd() string {
	return r._gmtCreateEnd
}

// SetGmtStatusUpdateEnd is GmtStatusUpdateEnd Setter
// 状态变更时间结束
func (r *AliexpresslogisticsabnormalorderqueryAPIRequest) SetGmtStatusUpdateEnd(_gmtStatusUpdateEnd string) error {
	r._gmtStatusUpdateEnd = _gmtStatusUpdateEnd
	r.Set("gmt_status_update_end", _gmtStatusUpdateEnd)
	return nil
}

// GetGmtStatusUpdateEnd GmtStatusUpdateEnd Getter
func (r AliexpresslogisticsabnormalorderqueryAPIRequest) GetGmtStatusUpdateEnd() string {
	return r._gmtStatusUpdateEnd
}

// SetTradeOrderId is TradeOrderId Setter
// 交易订单号
func (r *AliexpresslogisticsabnormalorderqueryAPIRequest) SetTradeOrderId(_tradeOrderId int64) error {
	r._tradeOrderId = _tradeOrderId
	r.Set("trade_order_id", _tradeOrderId)
	return nil
}

// GetTradeOrderId TradeOrderId Getter
func (r AliexpresslogisticsabnormalorderqueryAPIRequest) GetTradeOrderId() int64 {
	return r._tradeOrderId
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AliexpresslogisticsabnormalorderqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpresslogisticsabnormalorderqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *AliexpresslogisticsabnormalorderqueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AliexpresslogisticsabnormalorderqueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
