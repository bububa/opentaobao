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
type TaobaoBaichuanOpenaccountLoginRequest struct {
    model.Params
    // name
    name   string
}

// 初始化TaobaoBaichuanOpenaccountLoginRequest对象
func NewTaobaoBaichuanOpenaccountLoginRequest() *TaobaoBaichuanOpenaccountLoginRequest{
    return &TaobaoBaichuanOpenaccountLoginRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountLoginRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.login"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountLoginRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOpenaccountLoginRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOpenaccountLoginRequest) GetName() string {
    return r.name
}
