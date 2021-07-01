package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川验证找回密码验证码 API请求
taobao.baichuan.openaccount.resetcode.check

百川验证找回密码验证码
*/
type TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest struct {
    model.Params
    // name
    _name   string
}

// 初始化TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest对象
func NewTaobaoBaichuanOpenaccountResetcodeCheckRequest() *TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest{
    return &TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.resetcode.check"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest) GetName() string {
    return r._name
}
