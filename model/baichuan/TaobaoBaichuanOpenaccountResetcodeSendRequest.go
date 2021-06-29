package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川发送找回密码验证码 API请求
taobao.baichuan.openaccount.resetcode.send

百川发送找回密码验证码
*/
type TaobaoBaichuanOpenaccountResetcodeSendRequest struct {
    model.Params
    // name
    _name   string
}

// 初始化TaobaoBaichuanOpenaccountResetcodeSendRequest对象
func NewTaobaoBaichuanOpenaccountResetcodeSendRequest() *TaobaoBaichuanOpenaccountResetcodeSendRequest{
    return &TaobaoBaichuanOpenaccountResetcodeSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountResetcodeSendRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.resetcode.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountResetcodeSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOpenaccountResetcodeSendRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOpenaccountResetcodeSendRequest) GetName() string {
    return r._name
}
