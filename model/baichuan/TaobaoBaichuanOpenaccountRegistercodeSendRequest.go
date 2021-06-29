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
type TaobaoBaichuanOpenaccountRegistercodeSendRequest struct {
    model.Params
    // name
    name   string
}

// 初始化TaobaoBaichuanOpenaccountRegistercodeSendRequest对象
func NewTaobaoBaichuanOpenaccountRegistercodeSendRequest() *TaobaoBaichuanOpenaccountRegistercodeSendRequest{
    return &TaobaoBaichuanOpenaccountRegistercodeSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountRegistercodeSendRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.registercode.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountRegistercodeSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOpenaccountRegistercodeSendRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOpenaccountRegistercodeSendRequest) GetName() string {
    return r.name
}
