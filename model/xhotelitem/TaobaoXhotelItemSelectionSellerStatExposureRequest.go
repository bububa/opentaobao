package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
选品曝光趋势 API请求
taobao.xhotel.item.selection.seller.stat.exposure

用于提供给商家获取选品曝光趋势
*/
type TaobaoXhotelItemSelectionSellerStatExposureRequest struct {
    model.Params
    // 日期 默认为昨天
    date   string
    // hid  默认为all
    hid   string
    // 默认为all
    vendor   string
    // 默认为all
    supplier   string
    // 酒店编码
    outHid   string
}

// 初始化TaobaoXhotelItemSelectionSellerStatExposureRequest对象
func NewTaobaoXhotelItemSelectionSellerStatExposureRequest() *TaobaoXhotelItemSelectionSellerStatExposureRequest{
    return &TaobaoXhotelItemSelectionSellerStatExposureRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelItemSelectionSellerStatExposureRequest) GetApiMethodName() string {
    return "taobao.xhotel.item.selection.seller.stat.exposure"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelItemSelectionSellerStatExposureRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Date Setter
// 日期 默认为昨天
func (r *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetDate(date string) error {
    r.date = date
    r.Set("date", date)
    return nil
}

// Date Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureRequest) GetDate() string {
    return r.date
}
// Hid Setter
// hid  默认为all
func (r *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetHid(hid string) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureRequest) GetHid() string {
    return r.hid
}
// Vendor Setter
// 默认为all
func (r *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureRequest) GetVendor() string {
    return r.vendor
}
// Supplier Setter
// 默认为all
func (r *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetSupplier(supplier string) error {
    r.supplier = supplier
    r.Set("supplier", supplier)
    return nil
}

// Supplier Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureRequest) GetSupplier() string {
    return r.supplier
}
// OutHid Setter
// 酒店编码
func (r *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetOutHid(outHid string) error {
    r.outHid = outHid
    r.Set("out_hid", outHid)
    return nil
}

// OutHid Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureRequest) GetOutHid() string {
    return r.outHid
}
