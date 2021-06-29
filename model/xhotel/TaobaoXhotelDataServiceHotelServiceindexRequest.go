package xhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店服务指数 API请求
taobao.xhotel.data.service.hotel.serviceindex

酒店服务指数
*/
type TaobaoXhotelDataServiceHotelServiceindexRequest struct {
    model.Params
    // 酒店id
    _hid   int64
    // 渠道商名称
    _vendor   string
    // 1
    _startRow   int64
    // 10
    _pageSize   int64
    // 查询时间段结束
    _reportEndDate   string
    // 查询时间段开始
    _reportStartDate   string
    // 供应商名称
    _supplier   string
}

// 初始化TaobaoXhotelDataServiceHotelServiceindexRequest对象
func NewTaobaoXhotelDataServiceHotelServiceindexRequest() *TaobaoXhotelDataServiceHotelServiceindexRequest{
    return &TaobaoXhotelDataServiceHotelServiceindexRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetApiMethodName() string {
    return "taobao.xhotel.data.service.hotel.serviceindex"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Hid Setter
// 酒店id
func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetHid() int64 {
    return r._hid
}
// Vendor Setter
// 渠道商名称
func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetVendor() string {
    return r._vendor
}
// StartRow Setter
// 1
func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetStartRow(_startRow int64) error {
    r._startRow = _startRow
    r.Set("start_row", _startRow)
    return nil
}

// StartRow Getter
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetStartRow() int64 {
    return r._startRow
}
// PageSize Setter
// 10
func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetPageSize() int64 {
    return r._pageSize
}
// ReportEndDate Setter
// 查询时间段结束
func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetReportEndDate(_reportEndDate string) error {
    r._reportEndDate = _reportEndDate
    r.Set("report_end_date", _reportEndDate)
    return nil
}

// ReportEndDate Getter
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetReportEndDate() string {
    return r._reportEndDate
}
// ReportStartDate Setter
// 查询时间段开始
func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetReportStartDate(_reportStartDate string) error {
    r._reportStartDate = _reportStartDate
    r.Set("report_start_date", _reportStartDate)
    return nil
}

// ReportStartDate Getter
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetReportStartDate() string {
    return r._reportStartDate
}
// Supplier Setter
// 供应商名称
func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetSupplier(_supplier string) error {
    r._supplier = _supplier
    r.Set("supplier", _supplier)
    return nil
}

// Supplier Getter
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetSupplier() string {
    return r._supplier
}
