package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川发送注册验证码 API请求
taobao.baichuan.openaccount.registercode.send

百川发送注册验证码
*/
type TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest struct {
    model.Params
    // name
    _name   string
}

// 初始化TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest对象
func NewTaobaoBaichuanOpenaccountRegistercodeSendRequest() *TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest{
    return &TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.registercode.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest) GetName() string {
    return r._name
}
