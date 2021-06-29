package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家更新库存 API请求
taobao.drug.quantity.update

商家通过top接口可以直接修改商品库存
*/
type TaobaoDrugQuantityUpdateRequest struct {
    model.Params
    // 外部店铺ID
    _outStoreId   string
    // 外部商品ID
    _outItemId   string
    // 库存数量
    _quantity   int64
}

// 初始化TaobaoDrugQuantityUpdateRequest对象
func NewTaobaoDrugQuantityUpdateRequest() *TaobaoDrugQuantityUpdateRequest{
    return &TaobaoDrugQuantityUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDrugQuantityUpdateRequest) GetApiMethodName() string {
    return "taobao.drug.quantity.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDrugQuantityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutStoreId Setter
// 外部店铺ID
func (r *TaobaoDrugQuantityUpdateRequest) SetOutStoreId(_outStoreId string) error {
    r._outStoreId = _outStoreId
    r.Set("out_store_id", _outStoreId)
    return nil
}

// OutStoreId Getter
func (r TaobaoDrugQuantityUpdateRequest) GetOutStoreId() string {
    return r._outStoreId
}
// OutItemId Setter
// 外部商品ID
func (r *TaobaoDrugQuantityUpdateRequest) SetOutItemId(_outItemId string) error {
    r._outItemId = _outItemId
    r.Set("out_item_id", _outItemId)
    return nil
}

// OutItemId Getter
func (r TaobaoDrugQuantityUpdateRequest) GetOutItemId() string {
    return r._outItemId
}
// Quantity Setter
// 库存数量
func (r *TaobaoDrugQuantityUpdateRequest) SetQuantity(_quantity int64) error {
    r._quantity = _quantity
    r.Set("quantity", _quantity)
    return nil
}

// Quantity Getter
func (r TaobaoDrugQuantityUpdateRequest) GetQuantity() int64 {
    return r._quantity
}
