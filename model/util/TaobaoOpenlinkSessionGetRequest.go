package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取授权session信息 APIRequest
taobao.openlink.session.get

帮助第三方isv生成三方session
*/
type TaobaoOpenlinkSessionGetRequest struct {
    model.Params

    // 授权码
    code   string 

}

func NewTaobaoOpenlinkSessionGetRequest() *TaobaoOpenlinkSessionGetRequest{
    return &TaobaoOpenlinkSessionGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenlinkSessionGetRequest) GetApiMethodName() string {
    return "taobao.openlink.session.get"
}

func (r TaobaoOpenlinkSessionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenlinkSessionGetRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r TaobaoOpenlinkSessionGetRequest) GetCode() string {
    return r.code
}

