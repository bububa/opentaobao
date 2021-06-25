package xhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店服务指数 APIRequest
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

func NewTaobaoXhotelDataServiceHotelServiceindexRequest() *TaobaoXhotelDataServiceHotelServiceindexRequest{
    return &TaobaoXhotelDataServiceHotelServiceindexRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetApiMethodName() string {
    return "taobao.xhotel.data.service.hotel.serviceindex"
}

func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetHid() int64 {
    return r.hid
}

func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetStartRow(startRow int64) error {
    r.startRow = startRow
    r.Set("start_row", startRow)
    return nil
}

func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetStartRow() int64 {
    return r.startRow
}

func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetReportEndDate(reportEndDate string) error {
    r.reportEndDate = reportEndDate
    r.Set("report_end_date", reportEndDate)
    return nil
}

func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetReportEndDate() string {
    return r.reportEndDate
}

func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetReportStartDate(reportStartDate string) error {
    r.reportStartDate = reportStartDate
    r.Set("report_start_date", reportStartDate)
    return nil
}

func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetReportStartDate() string {
    return r.reportStartDate
}

func (r *TaobaoXhotelDataServiceHotelServiceindexRequest) SetSupplier(supplier string) error {
    r.supplier = supplier
    r.Set("supplier", supplier)
    return nil
}

func (r TaobaoXhotelDataServiceHotelServiceindexRequest) GetSupplier() string {
    return r.supplier
}

