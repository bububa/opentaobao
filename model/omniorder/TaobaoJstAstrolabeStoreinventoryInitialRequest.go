package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
后端商品库存初始化 API请求
taobao.jst.astrolabe.storeinventory.initial

初始化电商仓或门店库存，该接口一次可以初始化多个门店(或电商仓)的多个商品的多种类型库存。此接口只能使用一次，后续所有的库存变动都需走增量库存同步接口。
*/
type TaobaoJstAstrolabeStoreinventoryInitialRequest struct {
    model.Params
    // 操作时间
    _operationTime   string
    // 门店列表
    _stores   []Store
}

// 初始化TaobaoJstAstrolabeStoreinventoryInitialRequest对象
func NewTaobaoJstAstrolabeStoreinventoryInitialRequest() *TaobaoJstAstrolabeStoreinventoryInitialRequest{
    return &TaobaoJstAstrolabeStoreinventoryInitialRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryInitialRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.storeinventory.initial"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryInitialRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OperationTime Setter
// 操作时间
func (r *TaobaoJstAstrolabeStoreinventoryInitialRequest) SetOperationTime(_operationTime string) error {
    r._operationTime = _operationTime
    r.Set("operation_time", _operationTime)
    return nil
}

// OperationTime Getter
func (r TaobaoJstAstrolabeStoreinventoryInitialRequest) GetOperationTime() string {
    return r._operationTime
}
// Stores Setter
// 门店列表
func (r *TaobaoJstAstrolabeStoreinventoryInitialRequest) SetStores(_stores []Store) error {
    r._stores = _stores
    r.Set("stores", _stores)
    return nil
}

// Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryInitialRequest) GetStores() []Store {
    return r._stores
}
