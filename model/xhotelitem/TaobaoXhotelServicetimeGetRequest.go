package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询实体对应的服务时间数据 API请求
taobao.xhotel.servicetime.get

通过实体来获取对应的服务时间数据
*/
type TaobaoXhotelServicetimeGetRequest struct {
    model.Params
    // 酒店id
    _hid   int64
}

// 初始化TaobaoXhotelServicetimeGetRequest对象
func NewTaobaoXhotelServicetimeGetRequest() *TaobaoXhotelServicetimeGetRequest{
    return &TaobaoXhotelServicetimeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelServicetimeGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.servicetime.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelServicetimeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Hid Setter
// 酒店id
func (r *TaobaoXhotelServicetimeGetRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelServicetimeGetRequest) GetHid() int64 {
    return r._hid
}
