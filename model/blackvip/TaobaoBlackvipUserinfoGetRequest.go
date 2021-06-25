package blackvip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
88VIP用户信息查询 APIRequest
taobao.blackvip.userinfo.get

查询88VIP用户信息，比如用户是否是88VIP，88VIP的失效时间等
*/
type TaobaoBlackvipUserinfoGetRequest struct {
    model.Params

}

func NewTaobaoBlackvipUserinfoGetRequest() *TaobaoBlackvipUserinfoGetRequest{
    return &TaobaoBlackvipUserinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBlackvipUserinfoGetRequest) GetApiMethodName() string {
    return "taobao.blackvip.userinfo.get"
}

func (r TaobaoBlackvipUserinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


