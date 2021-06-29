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
    hid   int64
    // 渠道商名称
    vendor   string
    // 1
    startRow   int64
    // 10
    pageSize   int64
    // 查询时间段结束
    reportEndDate   string
    // 查询时间段开始
    reportStartDate   string
    // 供应商名称
    supplier   string
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
func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetHid() int64 {
    return r.hid
}
// Vendor Setter
// 渠道商名称
func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetVendor() string {
    return r.vendor
}
// StartRow Setter
// 1
func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetStartRow(startRow int64) error {
    r.startRow = startRow
    r.Set("start_row", startRow)
    return nil
}

// StartRow Getter
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetStartRow() int64 {
    return r.startRow
}
// PageSize Setter
// 10
func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetPageSize() int64 {
    return r.pageSize
}
// ReportEndDate Setter
// 查询时间段结束
func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetReportEndDate(reportEndDate string) error {
    r.reportEndDate = reportEndDate
    r.Set("report_end_date", reportEndDate)
    return nil
}

// ReportEndDate Getter
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetReportEndDate() string {
    return r.reportEndDate
}
// ReportStartDate Setter
// 查询时间段开始
func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetReportStartDate(reportStartDate string) error {
    r.reportStartDate = reportStartDate
    r.Set("report_start_date", reportStartDate)
    return nil
}

// ReportStartDate Getter
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetReportStartDate() string {
    return r.reportStartDate
}
// Supplier Setter
// 供应商名称
func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetSupplier(supplier string) error {
    r.supplier = supplier
    r.Set("supplier", supplier)
    return nil
}

// Supplier Getter
func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetSupplier() string {
    return r.supplier
}
