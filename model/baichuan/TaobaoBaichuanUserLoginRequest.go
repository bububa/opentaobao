package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川H5登录 API请求
taobao.baichuan.user.login

百川H5登录
*/
type TaobaoBaichuanUserLoginRequest struct {
    model.Params
    // name
    name   string
}

// 初始化TaobaoBaichuanUserLoginRequest对象
func NewTaobaoBaichuanUserLoginRequest() *TaobaoBaichuanUserLoginRequest{
    return &TaobaoBaichuanUserLoginRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanUserLoginRequest) GetApiMethodName() string {
    return "taobao.baichuan.user.login"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanUserLoginRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanUserLoginRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanUserLoginRequest) GetName() string {
    return r.name
}
