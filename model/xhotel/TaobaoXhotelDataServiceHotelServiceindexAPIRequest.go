package xhotel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelDataServiceHotelServiceindexAPIRequest 酒店服务指数 API请求
// taobao.xhotel.data.service.hotel.serviceindex
//
// 酒店服务指数
type TaobaoXhotelDataServiceHotelServiceindexAPIRequest struct {
	model.Params
	// 渠道商名称
	_vendor string
	// 查询时间段结束
	_reportEndDate string
	// 查询时间段开始
	_reportStartDate string
	// 供应商名称
	_supplier string
	// 酒店id
	_hid int64
	// 1
	_startRow int64
	// 10
	_pageSize int64
}

// NewTaobaoXhotelDataServiceHotelServiceindexRequest 初始化TaobaoXhotelDataServiceHotelServiceindexAPIRequest对象
func NewTaobaoXhotelDataServiceHotelServiceindexRequest() *TaobaoXhotelDataServiceHotelServiceindexAPIRequest {
	return &TaobaoXhotelDataServiceHotelServiceindexAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelDataServiceHotelServiceindexAPIRequest) Reset() {
	r._vendor = ""
	r._reportEndDate = ""
	r._reportStartDate = ""
	r._supplier = ""
	r._hid = 0
	r._startRow = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelDataServiceHotelServiceindexAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.data.service.hotel.serviceindex"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelDataServiceHotelServiceindexAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelDataServiceHotelServiceindexAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 渠道商名称
func (r *TaobaoXhotelDataServiceHotelServiceindexAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelDataServiceHotelServiceindexAPIRequest) GetVendor() string {
	return r._vendor
}

// SetReportEndDate is ReportEndDate Setter
// 查询时间段结束
func (r *TaobaoXhotelDataServiceHotelServiceindexAPIRequest) SetReportEndDate(_reportEndDate string) error {
	r._reportEndDate = _reportEndDate
	r.Set("report_end_date", _reportEndDate)
	return nil
}

// GetReportEndDate ReportEndDate Getter
func (r TaobaoXhotelDataServiceHotelServiceindexAPIRequest) GetReportEndDate() string {
	return r._reportEndDate
}

// SetReportStartDate is ReportStartDate Setter
// 查询时间段开始
func (r *TaobaoXhotelDataServiceHotelServiceindexAPIRequest) SetReportStartDate(_reportStartDate string) error {
	r._reportStartDate = _reportStartDate
	r.Set("report_start_date", _reportStartDate)
	return nil
}

// GetReportStartDate ReportStartDate Getter
func (r TaobaoXhotelDataServiceHotelServiceindexAPIRequest) GetReportStartDate() string {
	return r._reportStartDate
}

// SetSupplier is Supplier Setter
// 供应商名称
func (r *TaobaoXhotelDataServiceHotelServiceindexAPIRequest) SetSupplier(_supplier string) error {
	r._supplier = _supplier
	r.Set("supplier", _supplier)
	return nil
}

// GetSupplier Supplier Getter
func (r TaobaoXhotelDataServiceHotelServiceindexAPIRequest) GetSupplier() string {
	return r._supplier
}

// SetHid is Hid Setter
// 酒店id
func (r *TaobaoXhotelDataServiceHotelServiceindexAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoXhotelDataServiceHotelServiceindexAPIRequest) GetHid() int64 {
	return r._hid
}

// SetStartRow is StartRow Setter
// 1
func (r *TaobaoXhotelDataServiceHotelServiceindexAPIRequest) SetStartRow(_startRow int64) error {
	r._startRow = _startRow
	r.Set("start_row", _startRow)
	return nil
}

// GetStartRow StartRow Getter
func (r TaobaoXhotelDataServiceHotelServiceindexAPIRequest) GetStartRow() int64 {
	return r._startRow
}

// SetPageSize is PageSize Setter
// 10
func (r *TaobaoXhotelDataServiceHotelServiceindexAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoXhotelDataServiceHotelServiceindexAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoXhotelDataServiceHotelServiceindexAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelDataServiceHotelServiceindexRequest()
	},
}

// GetTaobaoXhotelDataServiceHotelServiceindexRequest 从 sync.Pool 获取 TaobaoXhotelDataServiceHotelServiceindexAPIRequest
func GetTaobaoXhotelDataServiceHotelServiceindexAPIRequest() *TaobaoXhotelDataServiceHotelServiceindexAPIRequest {
	return poolTaobaoXhotelDataServiceHotelServiceindexAPIRequest.Get().(*TaobaoXhotelDataServiceHotelServiceindexAPIRequest)
}

// ReleaseTaobaoXhotelDataServiceHotelServiceindexAPIRequest 将 TaobaoXhotelDataServiceHotelServiceindexAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelDataServiceHotelServiceindexAPIRequest(v *TaobaoXhotelDataServiceHotelServiceindexAPIRequest) {
	v.Reset()
	poolTaobaoXhotelDataServiceHotelServiceindexAPIRequest.Put(v)
}
