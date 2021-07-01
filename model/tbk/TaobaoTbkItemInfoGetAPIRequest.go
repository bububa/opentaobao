package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-淘宝客商品详情查询(简版) API请求
taobao.tbk.item.info.get

淘宝客商品详情查询（简版）
*/
type TaobaoTbkItemInfoGetAPIRequest struct {
    model.Params
    // 商品ID串，用,分割，最大40个
    _numIids   string
    // 链接形式：1：PC，2：无线，默认：１
    _platform   int64
    // ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
    _ip   string
}

// 初始化TaobaoTbkItemInfoGetAPIRequest对象
func NewTaobaoTbkItemInfoGetRequest() *TaobaoTbkItemInfoGetAPIRequest{
    return &TaobaoTbkItemInfoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkItemInfoGetAPIRequest) GetApiMethodName() string {
    return "taobao.tbk.item.info.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkItemInfoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NumIids Setter
// 商品ID串，用,分割，最大40个
func (r *TaobaoTbkItemInfoGetAPIRequest) SetNumIids(_numIids string) error {
    r._numIids = _numIids
    r.Set("num_iids", _numIids)
    return nil
}

// NumIids Getter
func (r TaobaoTbkItemInfoGetAPIRequest) GetNumIids() string {
    return r._numIids
}
// Platform Setter
// 链接形式：1：PC，2：无线，默认：１
func (r *TaobaoTbkItemInfoGetAPIRequest) SetPlatform(_platform int64) error {
    r._platform = _platform
    r.Set("platform", _platform)
    return nil
}

// Platform Getter
func (r TaobaoTbkItemInfoGetAPIRequest) GetPlatform() int64 {
    return r._platform
}
// Ip Setter
// ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
func (r *TaobaoTbkItemInfoGetAPIRequest) SetIp(_ip string) error {
    r._ip = _ip
    r.Set("ip", _ip)
    return nil
}

// Ip Getter
func (r TaobaoTbkItemInfoGetAPIRequest) GetIp() string {
    return r._ip
}
