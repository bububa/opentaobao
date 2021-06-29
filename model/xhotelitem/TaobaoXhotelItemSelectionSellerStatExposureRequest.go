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
    _date   string
    // hid  默认为all
    _hid   string
    // 默认为all
    _vendor   string
    // 默认为all
    _supplier   string
    // 酒店编码
    _outHid   string
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
func (r *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetDate(_date string) error {
    r._date = _date
    r.Set("date", _date)
    return nil
}

// Date Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureRequest) GetDate() string {
    return r._date
}
// Hid Setter
// hid  默认为all
func (r *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetHid(_hid string) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureRequest) GetHid() string {
    return r._hid
}
// Vendor Setter
// 默认为all
func (r *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureRequest) GetVendor() string {
    return r._vendor
}
// Supplier Setter
// 默认为all
func (r *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetSupplier(_supplier string) error {
    r._supplier = _supplier
    r.Set("supplier", _supplier)
    return nil
}

// Supplier Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureRequest) GetSupplier() string {
    return r._supplier
}
// OutHid Setter
// 酒店编码
func (r *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetOutHid(_outHid string) error {
    r._outHid = _outHid
    r.Set("out_hid", _outHid)
    return nil
}

// OutHid Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureRequest) GetOutHid() string {
    return r._outHid
}
