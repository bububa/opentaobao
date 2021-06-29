package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川登录二次验证 API请求
taobao.baichuan.openaccount.logindoublecheck

百川登录二次验证
*/
type TaobaoBaichuanOpenaccountLogindoublecheckRequest struct {
    model.Params
    // name
    _name   string
}

// 初始化TaobaoBaichuanOpenaccountLogindoublecheckRequest对象
func NewTaobaoBaichuanOpenaccountLogindoublecheckRequest() *TaobaoBaichuanOpenaccountLogindoublecheckRequest{
    return &TaobaoBaichuanOpenaccountLogindoublecheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountLogindoublecheckRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.logindoublecheck"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountLogindoublecheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOpenaccountLogindoublecheckRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOpenaccountLogindoublecheckRequest) GetName() string {
    return r._name
}
