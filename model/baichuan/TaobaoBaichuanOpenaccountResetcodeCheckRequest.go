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
type TaobaoBaichuanOpenaccountResetcodeCheckRequest struct {
    model.Params
    // name
    name   string
}

// 初始化TaobaoBaichuanOpenaccountResetcodeCheckRequest对象
func NewTaobaoBaichuanOpenaccountResetcodeCheckRequest() *TaobaoBaichuanOpenaccountResetcodeCheckRequest{
    return &TaobaoBaichuanOpenaccountResetcodeCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountResetcodeCheckRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.resetcode.check"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountResetcodeCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOpenaccountResetcodeCheckRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOpenaccountResetcodeCheckRequest) GetName() string {
    return r.name
}
