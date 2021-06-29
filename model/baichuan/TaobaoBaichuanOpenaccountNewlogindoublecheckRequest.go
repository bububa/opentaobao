package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川新登录二次验证 API请求
taobao.baichuan.openaccount.newlogindoublecheck

百川新登录二次验证
*/
type TaobaoBaichuanOpenaccountNewlogindoublecheckRequest struct {
    model.Params
    // name
    _name   string
}

// 初始化TaobaoBaichuanOpenaccountNewlogindoublecheckRequest对象
func NewTaobaoBaichuanOpenaccountNewlogindoublecheckRequest() *TaobaoBaichuanOpenaccountNewlogindoublecheckRequest{
    return &TaobaoBaichuanOpenaccountNewlogindoublecheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountNewlogindoublecheckRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.newlogindoublecheck"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountNewlogindoublecheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOpenaccountNewlogindoublecheckRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOpenaccountNewlogindoublecheckRequest) GetName() string {
    return r._name
}
