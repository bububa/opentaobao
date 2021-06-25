package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取饿了么用户信息 APIRequest
taobao.miniapp.eleuserinfo.get

获取饿了么用户信息
*/
type TaobaoMiniappEleuserinfoGetRequest struct {
    model.Params

    // 怪兽业务方
    bizScence   string 

}

func NewTaobaoMiniappEleuserinfoGetRequest() *TaobaoMiniappEleuserinfoGetRequest{
    return &TaobaoMiniappEleuserinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappEleuserinfoGetRequest) GetApiMethodName() string {
    return "taobao.miniapp.eleuserinfo.get"
}

func (r TaobaoMiniappEleuserinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappEleuserinfoGetRequest) SetBizScence(bizScence string) error {
    r.bizScence = bizScence
    r.Set("biz_scence", bizScence)
    return nil
}

func (r TaobaoMiniappEleuserinfoGetRequest) GetBizScence() string {
    return r.bizScence
}

