package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据卖家id，获取关联的所有包 API请求
taobao.vmarket.eticket.package.base.list.get

根据卖家id，获取关联的所有包
*/
type TaobaoVmarketEticketPackageBaseListGetRequest struct {
    model.Params
}

// 初始化TaobaoVmarketEticketPackageBaseListGetRequest对象
func NewTaobaoVmarketEticketPackageBaseListGetRequest() *TaobaoVmarketEticketPackageBaseListGetRequest{
    return &TaobaoVmarketEticketPackageBaseListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketPackageBaseListGetRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.package.base.list.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketPackageBaseListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
