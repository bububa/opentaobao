package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川检查注册验证码 API请求
taobao.baichuan.openaccount.registercode.check

百川检查注册验证码
*/
type TaobaoBaichuanOpenaccountRegistercodeCheckRequest struct {
    model.Params
    // name
    name   string
}

// 初始化TaobaoBaichuanOpenaccountRegistercodeCheckRequest对象
func NewTaobaoBaichuanOpenaccountRegistercodeCheckRequest() *TaobaoBaichuanOpenaccountRegistercodeCheckRequest{
    return &TaobaoBaichuanOpenaccountRegistercodeCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountRegistercodeCheckRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.registercode.check"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountRegistercodeCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOpenaccountRegistercodeCheckRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOpenaccountRegistercodeCheckRequest) GetName() string {
    return r.name
}
