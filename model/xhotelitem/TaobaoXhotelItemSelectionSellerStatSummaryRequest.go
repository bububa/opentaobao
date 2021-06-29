package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家数据-选品整体概况 API请求
taobao.xhotel.item.selection.seller.stat.summary

商家数据-选品整体概况
*/
type TaobaoXhotelItemSelectionSellerStatSummaryRequest struct {
    model.Params
    // vendor 默认为all
    vendor   string
    // 日期  默认为昨天
    date   string
    // hid  默认为all
    hid   string
    // supplier 默认为all
    supplier   string
    // 外部酒店编码
    outHid   string
}

// 初始化TaobaoXhotelItemSelectionSellerStatSummaryRequest对象
func NewTaobaoXhotelItemSelectionSellerStatSummaryRequest() *TaobaoXhotelItemSelectionSellerStatSummaryRequest{
    return &TaobaoXhotelItemSelectionSellerStatSummaryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelItemSelectionSellerStatSummaryRequest) GetApiMethodName() string {
    return "taobao.xhotel.item.selection.seller.stat.summary"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelItemSelectionSellerStatSummaryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Vendor Setter
// vendor 默认为all
func (r *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryRequest) GetVendor() string {
    return r.vendor
}
// Date Setter
// 日期  默认为昨天
func (r *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetDate(date string) error {
    r.date = date
    r.Set("date", date)
    return nil
}

// Date Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryRequest) GetDate() string {
    return r.date
}
// Hid Setter
// hid  默认为all
func (r *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetHid(hid string) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryRequest) GetHid() string {
    return r.hid
}
// Supplier Setter
// supplier 默认为all
func (r *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetSupplier(supplier string) error {
    r.supplier = supplier
    r.Set("supplier", supplier)
    return nil
}

// Supplier Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryRequest) GetSupplier() string {
    return r.supplier
}
// OutHid Setter
// 外部酒店编码
func (r *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetOutHid(outHid string) error {
    r.outHid = outHid
    r.Set("out_hid", outHid)
    return nil
}

// OutHid Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryRequest) GetOutHid() string {
    return r.outHid
}
