package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存查询接口 APIRequest
taobao.jst.astrolabe.storeinventory.itemquery

查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
*/
type TaobaoJstAstrolabeStoreinventoryItemqueryRequest struct {
    model.Params

    // 门店信息
    stores   []Store 

}

func NewTaobaoJstAstrolabeStoreinventoryItemqueryRequest() *TaobaoJstAstrolabeStoreinventoryItemqueryRequest{
    return &TaobaoJstAstrolabeStoreinventoryItemqueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstAstrolabeStoreinventoryItemqueryRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.storeinventory.itemquery"
}

func (r TaobaoJstAstrolabeStoreinventoryItemqueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstAstrolabeStoreinventoryItemqueryRequest) SetStores(stores []Store) error {
    r.stores = stores
    r.Set("stores", stores)
    return nil
}

func (r TaobaoJstAstrolabeStoreinventoryItemqueryRequest) GetStores() []Store {
    return r.stores
}

