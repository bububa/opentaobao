package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川订单详情 API请求
taobao.baichuan.orderurl.get

百川订单详情
*/
type TaobaoBaichuanOrderurlGetAPIRequest struct {
    model.Params
    // name
    _name   string
}

// 初始化TaobaoBaichuanOrderurlGetAPIRequest对象
func NewTaobaoBaichuanOrderurlGetRequest() *TaobaoBaichuanOrderurlGetAPIRequest{
    return &TaobaoBaichuanOrderurlGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOrderurlGetAPIRequest) GetApiMethodName() string {
    return "taobao.baichuan.orderurl.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOrderurlGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOrderurlGetAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOrderurlGetAPIRequest) GetName() string {
    return r._name
}
