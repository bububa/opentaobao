package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
后端商品库存查询接口 API请求
taobao.jst.astrolabe.storeinventory.query

查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
*/
type TaobaoJstAstrolabeStoreinventoryQueryRequest struct {
    model.Params
    // 门店
    _stores   []Store
}

// 初始化TaobaoJstAstrolabeStoreinventoryQueryRequest对象
func NewTaobaoJstAstrolabeStoreinventoryQueryRequest() *TaobaoJstAstrolabeStoreinventoryQueryRequest{
    return &TaobaoJstAstrolabeStoreinventoryQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryQueryRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.storeinventory.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Stores Setter
// 门店
func (r *TaobaoJstAstrolabeStoreinventoryQueryRequest) SetStores(_stores []Store) error {
    r._stores = _stores
    r.Set("stores", _stores)
    return nil
}

// Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryQueryRequest) GetStores() []Store {
    return r._stores
}
