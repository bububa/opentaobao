package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除酒店接口 API请求
taobao.xhotel.delete

删除飞猪酒店数据接口
*/
type TaobaoXhotelDeleteAPIRequest struct {
    model.Params
    // 酒店id，传参方式  hid   或者   outer_id+vendor
    _hid   int64
    // 酒店vendor
    _vendor   string
    // 酒店编码
    _outerId   string
}

// 初始化TaobaoXhotelDeleteAPIRequest对象
func NewTaobaoXhotelDeleteRequest() *TaobaoXhotelDeleteAPIRequest{
    return &TaobaoXhotelDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelDeleteAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Hid Setter
// 酒店id，传参方式  hid   或者   outer_id+vendor
func (r *TaobaoXhotelDeleteAPIRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelDeleteAPIRequest) GetHid() int64 {
    return r._hid
}
// Vendor Setter
// 酒店vendor
func (r *TaobaoXhotelDeleteAPIRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelDeleteAPIRequest) GetVendor() string {
    return r._vendor
}
// OuterId Setter
// 酒店编码
func (r *TaobaoXhotelDeleteAPIRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelDeleteAPIRequest) GetOuterId() string {
    return r._outerId
}
