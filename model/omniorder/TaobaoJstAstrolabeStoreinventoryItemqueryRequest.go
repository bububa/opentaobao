package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存查询接口 API请求
taobao.jst.astrolabe.storeinventory.itemquery

查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
*/
type TaobaoJstAstrolabeStoreinventoryItemqueryRequest struct {
    model.Params
    // 门店信息
    _stores   []Store
}

// 初始化TaobaoJstAstrolabeStoreinventoryItemqueryRequest对象
func NewTaobaoJstAstrolabeStoreinventoryItemqueryRequest() *TaobaoJstAstrolabeStoreinventoryItemqueryRequest{
    return &TaobaoJstAstrolabeStoreinventoryItemqueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryItemqueryRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.storeinventory.itemquery"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryItemqueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Stores Setter
// 门店信息
func (r *TaobaoJstAstrolabeStoreinventoryItemqueryRequest) SetStores(_stores []Store) error {
    r._stores = _stores
    r.Set("stores", _stores)
    return nil
}

// Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryItemqueryRequest) GetStores() []Store {
    return r._stores
}
