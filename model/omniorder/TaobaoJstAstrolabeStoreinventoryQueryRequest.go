package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
后端商品库存查询接口 APIRequest
taobao.jst.astrolabe.storeinventory.query

查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
*/
type TaobaoJstAstrolabeStoreinventoryQueryRequest struct {
    model.Params

    // 门店
    stores   []Store 

}

func NewTaobaoJstAstrolabeStoreinventoryQueryRequest() *TaobaoJstAstrolabeStoreinventoryQueryRequest{
    return &TaobaoJstAstrolabeStoreinventoryQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstAstrolabeStoreinventoryQueryRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.storeinventory.query"
}

func (r TaobaoJstAstrolabeStoreinventoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstAstrolabeStoreinventoryQueryRequest) SetStores(stores []Store) error {
    r.stores = stores
    r.Set("stores", stores)
    return nil
}

func (r TaobaoJstAstrolabeStoreinventoryQueryRequest) GetStores() []Store {
    return r.stores
}

