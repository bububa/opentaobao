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
    _date   string
    // 酒店id  默认all
    _hid   string
    // vendor  默认all
    _vendor   string
    // supplier  默认all
    _supplier   string
    // 酒店编码
    _outHid   string
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
func (r *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetDate(_date string) error {
    r._date = _date
    r.Set("date", _date)
    return nil
}

// Date Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidRequest) GetDate() string {
    return r._date
}
// Hid Setter
// 酒店id  默认all
func (r *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetHid(_hid string) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidRequest) GetHid() string {
    return r._hid
}
// Vendor Setter
// vendor  默认all
func (r *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidRequest) GetVendor() string {
    return r._vendor
}
// Supplier Setter
// supplier  默认all
func (r *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetSupplier(_supplier string) error {
    r._supplier = _supplier
    r.Set("supplier", _supplier)
    return nil
}

// Supplier Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidRequest) GetSupplier() string {
    return r._supplier
}
// OutHid Setter
// 酒店编码
func (r *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetOutHid(_outHid string) error {
    r._outHid = _outHid
    r.Set("out_hid", _outHid)
    return nil
}

// OutHid Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidRequest) GetOutHid() string {
    return r._outHid
}
