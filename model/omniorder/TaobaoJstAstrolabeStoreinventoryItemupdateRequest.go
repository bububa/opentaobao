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
type TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest struct {
    model.Params
    // 门店列表
    _stores   []Store
    // 操作时间
    _operationTime   string
}

// 初始化TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest对象
func NewTaobaoJstAstrolabeStoreinventoryItemupdateRequest() *TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest{
    return &TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.storeinventory.itemupdate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Stores Setter
// 门店列表
func (r *TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) SetStores(_stores []Store) error {
    r._stores = _stores
    r.Set("stores", _stores)
    return nil
}

// Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) GetStores() []Store {
    return r._stores
}
// OperationTime Setter
// 操作时间
func (r *TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) SetOperationTime(_operationTime string) error {
    r._operationTime = _operationTime
    r.Set("operation_time", _operationTime)
    return nil
}

// OperationTime Getter
func (r TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) GetOperationTime() string {
    return r._operationTime
}
