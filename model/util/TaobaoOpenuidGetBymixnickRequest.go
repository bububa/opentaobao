package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过mixnick转换openuid APIRequest
taobao.openuid.get.bymixnick

通过mixnick转换openuid
*/
type TaobaoOpenuidGetBymixnickRequest struct {
    model.Params

    // 无线类应用获取到的混淆的nick
    mixNick   string 

}

func NewTaobaoOpenuidGetBymixnickRequest() *TaobaoOpenuidGetBymixnickRequest{
    return &TaobaoOpenuidGetBymixnickRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenuidGetBymixnickRequest) GetApiMethodName() string {
    return "taobao.openuid.get.bymixnick"
}

func (r TaobaoOpenuidGetBymixnickRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenuidGetBymixnickRequest) SetMixNick(mixNick string) error {
    r.mixNick = mixNick
    r.Set("mix_nick", mixNick)
    return nil
}

func (r TaobaoOpenuidGetBymixnickRequest) GetMixNick() string {
    return r.mixNick
}

