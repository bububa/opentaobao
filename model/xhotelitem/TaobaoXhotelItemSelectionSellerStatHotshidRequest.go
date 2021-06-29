package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链选品热销标准酒店覆盖情况 API请求
taobao.xhotel.item.selection.seller.stat.hotshid

供应链选品热销标准酒店覆盖情况
*/
type TaobaoXhotelItemSelectionSellerStatHotshidRequest struct {
    model.Params
    // 日期  默认为昨天
    date   string
    // 酒店id  默认all
    hid   string
    // vendor  默认all
    vendor   string
    // supplier  默认all
    supplier   string
    // 酒店编码
    outHid   string
}

// 初始化TaobaoXhotelItemSelectionSellerStatHotshidRequest对象
func NewTaobaoXhotelItemSelectionSellerStatHotshidRequest() *TaobaoXhotelItemSelectionSellerStatHotshidRequest{
    return &TaobaoXhotelItemSelectionSellerStatHotshidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelItemSelectionSellerStatHotshidRequest) GetApiMethodName() string {
    return "taobao.xhotel.item.selection.seller.stat.hotshid"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelItemSelectionSellerStatHotshidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Date Setter
// 日期  默认为昨天
func (r *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetDate(date string) error {
    r.date = date
    r.Set("date", date)
    return nil
}

// Date Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidRequest) GetDate() string {
    return r.date
}
// Hid Setter
// 酒店id  默认all
func (r *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetHid(hid string) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidRequest) GetHid() string {
    return r.hid
}
// Vendor Setter
// vendor  默认all
func (r *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidRequest) GetVendor() string {
    return r.vendor
}
// Supplier Setter
// supplier  默认all
func (r *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetSupplier(supplier string) error {
    r.supplier = supplier
    r.Set("supplier", supplier)
    return nil
}

// Supplier Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidRequest) GetSupplier() string {
    return r.supplier
}
// OutHid Setter
// 酒店编码
func (r *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetOutHid(outHid string) error {
    r.outHid = outHid
    r.Set("out_hid", outHid)
    return nil
}

// OutHid Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidRequest) GetOutHid() string {
    return r.outHid
}
