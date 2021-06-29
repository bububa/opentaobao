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
type TaobaoBaichuanOpenaccountRegisterRequest struct {
    model.Params
    // name
    name   string
}

// 初始化TaobaoBaichuanOpenaccountRegisterRequest对象
func NewTaobaoBaichuanOpenaccountRegisterRequest() *TaobaoBaichuanOpenaccountRegisterRequest{
    return &TaobaoBaichuanOpenaccountRegisterRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountRegisterRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.register"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOpenaccountRegisterRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOpenaccountRegisterRequest) GetName() string {
    return r.name
}
