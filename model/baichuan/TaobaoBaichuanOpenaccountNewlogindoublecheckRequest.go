package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川新登录二次验证 APIRequest
taobao.baichuan.openaccount.newlogindoublecheck

百川新登录二次验证
*/
type TaobaoBaichuanOpenaccountNewlogindoublecheckRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanOpenaccountNewlogindoublecheckRequest() *TaobaoBaichuanOpenaccountNewlogindoublecheckRequest{
    return &TaobaoBaichuanOpenaccountNewlogindoublecheckRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanOpenaccountNewlogindoublecheckRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.newlogindoublecheck"
}

func (r TaobaoBaichuanOpenaccountNewlogindoublecheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanOpenaccountNewlogindoublecheckRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanOpenaccountNewlogindoublecheckRequest) GetName() string {
    return r.name
}

