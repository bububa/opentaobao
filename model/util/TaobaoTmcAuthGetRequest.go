package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
TMC授权token APIRequest
taobao.tmc.auth.get

TMC连接授权Token
*/
type TaobaoTmcAuthGetRequest struct {
    model.Params

    // tmc组名
    group   string 

}

func NewTaobaoTmcAuthGetRequest() *TaobaoTmcAuthGetRequest{
    return &TaobaoTmcAuthGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTmcAuthGetRequest) GetApiMethodName() string {
    return "taobao.tmc.auth.get"
}

func (r TaobaoTmcAuthGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTmcAuthGetRequest) SetGroup(group string) error {
    r.group = group
    r.Set("group", group)
    return nil
}

func (r TaobaoTmcAuthGetRequest) GetGroup() string {
    return r.group
}

