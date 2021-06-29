package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存增量更新接口 API请求
taobao.jst.astrolabe.storeinventory.itemupdate

ERP调用该接口，增量更新门店或电商仓库存，该接口一次可以同时增量更新多个门店的多个商品的非确定性库存。
*/
type TaobaoJstAstrolabeStoreinventoryItemupdateRequest struct {
    model.Params
    // 门店列表
    stores   []Store
    // 操作时间
    operationTime   string
}

// 初始化TaobaoJstAstrolabeStoreinventoryItemupdateRequest对象
func NewTaobaoJstAstrolabeStoreinventoryItemupdateRequest() *TaobaoJstAstrolabeStoreinventoryItemupdateRequest{
    return &TaobaoJstAstrolabeStoreinventoryItemupdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryItemupdateRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.storeinventory.itemupdate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryItemupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Stores Setter
// 门店列表
func (r *TaobaoJstAstrolabeStoreinventoryItemupdateRequest) SetStores(stores []Store) error {
    r.stores = stores
    r.Set("stores", stores)
    return nil
}

// Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryItemupdateRequest) GetStores() []Store {
    return r.stores
}
// OperationTime Setter
// 操作时间
func (r *TaobaoJstAstrolabeStoreinventoryItemupdateRequest) SetOperationTime(operationTime string) error {
    r.operationTime = operationTime
    r.Set("operation_time", operationTime)
    return nil
}

// OperationTime Getter
func (r TaobaoJstAstrolabeStoreinventoryItemupdateRequest) GetOperationTime() string {
    return r.operationTime
}
