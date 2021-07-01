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
type TaobaoVmarketEticketPackageBaseListGetAPIRequest struct {
    model.Params
}

// 初始化TaobaoVmarketEticketPackageBaseListGetAPIRequest对象
func NewTaobaoVmarketEticketPackageBaseListGetRequest() *TaobaoVmarketEticketPackageBaseListGetAPIRequest{
    return &TaobaoVmarketEticketPackageBaseListGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketPackageBaseListGetAPIRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.package.base.list.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketPackageBaseListGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
