package xhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家服务指数查询 API请求
taobao.xhotel.data.service.seller.serviceindex

卖家服务指数查询
*/
type TaobaoXhotelDataServiceSellerServiceindexRequest struct {
    model.Params
    // 渠道商名称
    _vendor   string
    // 分页参数
    _startRow   int64
    // 分页参数
    _pageSize   int64
    // 查询截止日期
    _reportEndDate   string
    // 查询开始日期
    _reportStartDate   string
}

// 初始化TaobaoXhotelDataServiceSellerServiceindexRequest对象
func NewTaobaoXhotelDataServiceSellerServiceindexRequest() *TaobaoXhotelDataServiceSellerServiceindexRequest{
    return &TaobaoXhotelDataServiceSellerServiceindexRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelDataServiceSellerServiceindexRequest) GetApiMethodName() string {
    return "taobao.xhotel.data.service.seller.serviceindex"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelDataServiceSellerServiceindexRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Vendor Setter
// 渠道商名称
func (r *TaobaoXhotelDataServiceSellerServiceindexRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelDataServiceSellerServiceindexRequest) GetVendor() string {
    return r._vendor
}
// StartRow Setter
// 分页参数
func (r *TaobaoXhotelDataServiceSellerServiceindexRequest) SetStartRow(_startRow int64) error {
    r._startRow = _startRow
    r.Set("start_row", _startRow)
    return nil
}

// StartRow Getter
func (r TaobaoXhotelDataServiceSellerServiceindexRequest) GetStartRow() int64 {
    return r._startRow
}
// PageSize Setter
// 分页参数
func (r *TaobaoXhotelDataServiceSellerServiceindexRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoXhotelDataServiceSellerServiceindexRequest) GetPageSize() int64 {
    return r._pageSize
}
// ReportEndDate Setter
// 查询截止日期
func (r *TaobaoXhotelDataServiceSellerServiceindexRequest) SetReportEndDate(_reportEndDate string) error {
    r._reportEndDate = _reportEndDate
    r.Set("report_end_date", _reportEndDate)
    return nil
}

// ReportEndDate Getter
func (r TaobaoXhotelDataServiceSellerServiceindexRequest) GetReportEndDate() string {
    return r._reportEndDate
}
// ReportStartDate Setter
// 查询开始日期
func (r *TaobaoXhotelDataServiceSellerServiceindexRequest) SetReportStartDate(_reportStartDate string) error {
    r._reportStartDate = _reportStartDate
    r.Set("report_start_date", _reportStartDate)
    return nil
}

// ReportStartDate Getter
func (r TaobaoXhotelDataServiceSellerServiceindexRequest) GetReportStartDate() string {
    return r._reportStartDate
}
