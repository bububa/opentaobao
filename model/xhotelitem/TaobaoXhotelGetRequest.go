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
type TaobaoXhotelGetRequest struct {
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

// 初始化TaobaoXhotelGetRequest对象
func NewTaobaoXhotelGetRequest() *TaobaoXhotelGetRequest{
    return &TaobaoXhotelGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Hid Setter
// 废弃，请使用outer_id
func (r *TaobaoXhotelGetRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelGetRequest) GetHid() int64 {
    return r._hid
}
// OuterId Setter
// 卖家系统中的酒店ID。
func (r *TaobaoXhotelGetRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelGetRequest) GetOuterId() string {
    return r._outerId
}
// Vendor Setter
// 系统商，一般不用填写，使用须申请
func (r *TaobaoXhotelGetRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelGetRequest) GetVendor() string {
    return r._vendor
}
// NeedSaleInfo Setter
// 是否需要在售状态(默认false不需要)
func (r *TaobaoXhotelGetRequest) SetNeedSaleInfo(_needSaleInfo bool) error {
    r._needSaleInfo = _needSaleInfo
    r.Set("need_sale_info", _needSaleInfo)
    return nil
}

// NeedSaleInfo Getter
func (r TaobaoXhotelGetRequest) GetNeedSaleInfo() bool {
    return r._needSaleInfo
}
