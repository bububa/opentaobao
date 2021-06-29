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
    name   string
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
func (r *TaobaoBaichuanOpenaccountNewlogindoublecheckRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOpenaccountNewlogindoublecheckRequest) GetName() string {
    return r.name
}
