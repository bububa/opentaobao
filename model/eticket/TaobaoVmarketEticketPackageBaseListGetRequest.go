package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据卖家id，获取关联的所有包 APIRequest
taobao.vmarket.eticket.package.base.list.get

根据卖家id，获取关联的所有包
*/
type TaobaoVmarketEticketPackageBaseListGetRequest struct {
    model.Params

}

func NewTaobaoVmarketEticketPackageBaseListGetRequest() *TaobaoVmarketEticketPackageBaseListGetRequest{
    return &TaobaoVmarketEticketPackageBaseListGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVmarketEticketPackageBaseListGetRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.package.base.list.get"
}

func (r TaobaoVmarketEticketPackageBaseListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


