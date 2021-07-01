package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川用户名密码登录 API请求
taobao.baichuan.openaccount.login

百川用户名密码登录
*/
type TaobaoBaichuanOpenaccountLoginAPIRequest struct {
    model.Params
    // name
    _name   string
}

// 初始化TaobaoBaichuanOpenaccountLoginAPIRequest对象
func NewTaobaoBaichuanOpenaccountLoginRequest() *TaobaoBaichuanOpenaccountLoginAPIRequest{
    return &TaobaoBaichuanOpenaccountLoginAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountLoginAPIRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.login"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountLoginAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOpenaccountLoginAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOpenaccountLoginAPIRequest) GetName() string {
    return r._name
}
