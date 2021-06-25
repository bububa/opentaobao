package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户售后服务模板 APIRequest
taobao.aftersale.get

查询用户设置的售后服务模板，仅返回标题和id
*/
type TaobaoAftersaleGetRequest struct {
    model.Params

}

func NewTaobaoAftersaleGetRequest() *TaobaoAftersaleGetRequest{
    return &TaobaoAftersaleGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAftersaleGetRequest) GetApiMethodName() string {
    return "taobao.aftersale.get"
}

func (r TaobaoAftersaleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


