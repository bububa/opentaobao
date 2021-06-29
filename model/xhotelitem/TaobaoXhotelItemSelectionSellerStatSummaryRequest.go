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
    _vendor   string
    // 日期  默认为昨天
    _date   string
    // hid  默认为all
    _hid   string
    // supplier 默认为all
    _supplier   string
    // 外部酒店编码
    _outHid   string
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
func (r *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryRequest) GetVendor() string {
    return r._vendor
}
// Date Setter
// 日期  默认为昨天
func (r *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetDate(_date string) error {
    r._date = _date
    r.Set("date", _date)
    return nil
}

// Date Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryRequest) GetDate() string {
    return r._date
}
// Hid Setter
// hid  默认为all
func (r *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetHid(_hid string) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryRequest) GetHid() string {
    return r._hid
}
// Supplier Setter
// supplier 默认为all
func (r *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetSupplier(_supplier string) error {
    r._supplier = _supplier
    r.Set("supplier", _supplier)
    return nil
}

// Supplier Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryRequest) GetSupplier() string {
    return r._supplier
}
// OutHid Setter
// 外部酒店编码
func (r *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetOutHid(_outHid string) error {
    r._outHid = _outHid
    r.Set("out_hid", _outHid)
    return nil
}

// OutHid Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryRequest) GetOutHid() string {
    return r._outHid
}
