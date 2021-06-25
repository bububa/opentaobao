package xhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家服务指数查询 APIRequest
taobao.xhotel.data.service.seller.serviceindex

卖家服务指数查询
*/
type TaobaoXhotelDataServiceSellerServiceindexRequest struct {
    model.Params

    // 渠道商名称
    vendor   string 

    // 分页参数
    startRow   int64 

    // 分页参数
    pageSize   int64 

    // 查询截止日期
    reportEndDate   string 

    // 查询开始日期
    reportStartDate   string 

}

func NewTaobaoXhotelDataServiceSellerServiceindexRequest() *TaobaoXhotelDataServiceSellerServiceindexRequest{
    return &TaobaoXhotelDataServiceSellerServiceindexRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelDataServiceSellerServiceindexRequest) GetApiMethodName() string {
    return "taobao.xhotel.data.service.seller.serviceindex"
}

func (r TaobaoXhotelDataServiceSellerServiceindexRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelDataServiceSellerServiceindexRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelDataServiceSellerServiceindexRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelDataServiceSellerServiceindexRequest) SetStartRow(startRow int64) error {
    r.startRow = startRow
    r.Set("start_row", startRow)
    return nil
}

func (r TaobaoXhotelDataServiceSellerServiceindexRequest) GetStartRow() int64 {
    return r.startRow
}

func (r *TaobaoXhotelDataServiceSellerServiceindexRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoXhotelDataServiceSellerServiceindexRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoXhotelDataServiceSellerServiceindexRequest) SetReportEndDate(reportEndDate string) error {
    r.reportEndDate = reportEndDate
    r.Set("report_end_date", reportEndDate)
    return nil
}

func (r TaobaoXhotelDataServiceSellerServiceindexRequest) GetReportEndDate() string {
    return r.reportEndDate
}

func (r *TaobaoXhotelDataServiceSellerServiceindexRequest) SetReportStartDate(reportStartDate string) error {
    r.reportStartDate = reportStartDate
    r.Set("report_start_date", reportStartDate)
    return nil
}

func (r TaobaoXhotelDataServiceSellerServiceindexRequest) GetReportStartDate() string {
    return r.reportStartDate
}

