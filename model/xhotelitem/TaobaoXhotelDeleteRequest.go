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
    hid   int64
    // 酒店vendor
    vendor   string
    // 酒店编码
    outerId   string
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
func (r *TaobaoXhotelDeleteRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelDeleteRequest) GetHid() int64 {
    return r.hid
}
// Vendor Setter
// 酒店vendor
func (r *TaobaoXhotelDeleteRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelDeleteRequest) GetVendor() string {
    return r.vendor
}
// OuterId Setter
// 酒店编码
func (r *TaobaoXhotelDeleteRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelDeleteRequest) GetOuterId() string {
    return r.outerId
}
