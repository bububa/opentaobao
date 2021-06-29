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
type TaobaoXhotelDeleteRequest struct {
    model.Params
    // 酒店id，传参方式  hid   或者   outer_id+vendor
    _hid   int64
    // 酒店vendor
    _vendor   string
    // 酒店编码
    _outerId   string
}

// 初始化TaobaoXhotelDeleteRequest对象
func NewTaobaoXhotelDeleteRequest() *TaobaoXhotelDeleteRequest{
    return &TaobaoXhotelDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelDeleteRequest) GetApiMethodName() string {
    return "taobao.xhotel.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Hid Setter
// 酒店id，传参方式  hid   或者   outer_id+vendor
func (r *TaobaoXhotelDeleteRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelDeleteRequest) GetHid() int64 {
    return r._hid
}
// Vendor Setter
// 酒店vendor
func (r *TaobaoXhotelDeleteRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelDeleteRequest) GetVendor() string {
    return r._vendor
}
// OuterId Setter
// 酒店编码
func (r *TaobaoXhotelDeleteRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelDeleteRequest) GetOuterId() string {
    return r._outerId
}
