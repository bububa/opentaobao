package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川手淘信任登录 API请求
taobao.baichuan.user.loginbytoken

百川手淘信任登录
*/
type TaobaoBaichuanUserLoginbytokenRequest struct {
    model.Params
    // name
    _name   string
}

// 初始化TaobaoBaichuanUserLoginbytokenRequest对象
func NewTaobaoBaichuanUserLoginbytokenRequest() *TaobaoBaichuanUserLoginbytokenRequest{
    return &TaobaoBaichuanUserLoginbytokenRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanUserLoginbytokenRequest) GetApiMethodName() string {
    return "taobao.baichuan.user.loginbytoken"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanUserLoginbytokenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanUserLoginbytokenRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanUserLoginbytokenRequest) GetName() string {
    return r._name
}
