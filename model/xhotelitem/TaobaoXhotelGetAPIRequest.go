package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店查询接口 API请求
taobao.xhotel.get

酒店查询接口
*/
type TaobaoXhotelGetAPIRequest struct {
    model.Params
    // 废弃，请使用outer_id
    _hid   int64
    // 卖家系统中的酒店ID。
    _outerId   string
    // 系统商，一般不用填写，使用须申请
    _vendor   string
    // 是否需要在售状态(默认false不需要)
    _needSaleInfo   bool
}

// 初始化TaobaoXhotelGetAPIRequest对象
func NewTaobaoXhotelGetRequest() *TaobaoXhotelGetAPIRequest{
    return &TaobaoXhotelGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelGetAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Hid Setter
// 废弃，请使用outer_id
func (r *TaobaoXhotelGetAPIRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelGetAPIRequest) GetHid() int64 {
    return r._hid
}
// OuterId Setter
// 卖家系统中的酒店ID。
func (r *TaobaoXhotelGetAPIRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelGetAPIRequest) GetOuterId() string {
    return r._outerId
}
// Vendor Setter
// 系统商，一般不用填写，使用须申请
func (r *TaobaoXhotelGetAPIRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelGetAPIRequest) GetVendor() string {
    return r._vendor
}
// NeedSaleInfo Setter
// 是否需要在售状态(默认false不需要)
func (r *TaobaoXhotelGetAPIRequest) SetNeedSaleInfo(_needSaleInfo bool) error {
    r._needSaleInfo = _needSaleInfo
    r.Set("need_sale_info", _needSaleInfo)
    return nil
}

// NeedSaleInfo Getter
func (r TaobaoXhotelGetAPIRequest) GetNeedSaleInfo() bool {
    return r._needSaleInfo
}
