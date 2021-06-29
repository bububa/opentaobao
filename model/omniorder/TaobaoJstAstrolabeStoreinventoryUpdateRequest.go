package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
后端商品库存增量更新接口 API请求
taobao.jst.astrolabe.storeinventory.update

增量更新门店或电商仓库存，该接口一次可以同时增量更新多个门店的多个商品的非确定性库存
*/
type TaobaoJstAstrolabeStoreinventoryUpdateRequest struct {
    model.Params
    // 操作时间
    operationTime   string
    // 门店列表
    stores   []Store
}

// 初始化TaobaoJstAstrolabeStoreinventoryUpdateRequest对象
func NewTaobaoJstAstrolabeStoreinventoryUpdateRequest() *TaobaoJstAstrolabeStoreinventoryUpdateRequest{
    return &TaobaoJstAstrolabeStoreinventoryUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryUpdateRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.storeinventory.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OperationTime Setter
// 操作时间
func (r *TaobaoJstAstrolabeStoreinventoryUpdateRequest) SetOperationTime(operationTime string) error {
    r.operationTime = operationTime
    r.Set("operation_time", operationTime)
    return nil
}

// OperationTime Getter
func (r TaobaoJstAstrolabeStoreinventoryUpdateRequest) GetOperationTime() string {
    return r.operationTime
}
// Stores Setter
// 门店列表
func (r *TaobaoJstAstrolabeStoreinventoryUpdateRequest) SetStores(stores []Store) error {
    r.stores = stores
    r.Set("stores", stores)
    return nil
}

// Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryUpdateRequest) GetStores() []Store {
    return r.stores
}
