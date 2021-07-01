package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川账号注册 API请求
taobao.baichuan.openaccount.register

百川账号注册
*/
type TaobaoBaichuanOpenaccountRegisterAPIRequest struct {
    model.Params
    // name
    _name   string
}

// 初始化TaobaoBaichuanOpenaccountRegisterAPIRequest对象
func NewTaobaoBaichuanOpenaccountRegisterRequest() *TaobaoBaichuanOpenaccountRegisterAPIRequest{
    return &TaobaoBaichuanOpenaccountRegisterAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountRegisterAPIRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.register"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountRegisterAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOpenaccountRegisterAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOpenaccountRegisterAPIRequest) GetName() string {
    return r._name
}
