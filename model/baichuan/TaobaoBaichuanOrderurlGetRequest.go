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
type TaobaoBaichuanOrderurlGetRequest struct {
    model.Params
    // name
    name   string
}

// 初始化TaobaoBaichuanOrderurlGetRequest对象
func NewTaobaoBaichuanOrderurlGetRequest() *TaobaoBaichuanOrderurlGetRequest{
    return &TaobaoBaichuanOrderurlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOrderurlGetRequest) GetApiMethodName() string {
    return "taobao.baichuan.orderurl.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOrderurlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOrderurlGetRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOrderurlGetRequest) GetName() string {
    return r.name
}
