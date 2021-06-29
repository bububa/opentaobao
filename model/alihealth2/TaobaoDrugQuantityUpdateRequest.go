package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家更新库存 APIRequest
taobao.drug.quantity.update

商家通过top接口可以直接修改商品库存
*/
type TaobaoDrugQuantityUpdateRequest struct {
    model.Params

    // 外部店铺ID
    outStoreId   string 

    // 外部商品ID
    outItemId   string 

    // 库存数量
    quantity   int64 

}

func NewTaobaoDrugQuantityUpdateRequest() *TaobaoDrugQuantityUpdateRequest{
    return &TaobaoDrugQuantityUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDrugQuantityUpdateRequest) GetApiMethodName() string {
    return "taobao.drug.quantity.update"
}

func (r TaobaoDrugQuantityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDrugQuantityUpdateRequest) SetOutStoreId(outStoreId string) error {
    r.outStoreId = outStoreId
    r.Set("out_store_id", outStoreId)
    return nil
}

func (r TaobaoDrugQuantityUpdateRequest) GetOutStoreId() string {
    return r.outStoreId
}

func (r *TaobaoDrugQuantityUpdateRequest) SetOutItemId(outItemId string) error {
    r.outItemId = outItemId
    r.Set("out_item_id", outItemId)
    return nil
}

func (r TaobaoDrugQuantityUpdateRequest) GetOutItemId() string {
    return r.outItemId
}

func (r *TaobaoDrugQuantityUpdateRequest) SetQuantity(quantity int64) error {
    r.quantity = quantity
    r.Set("quantity", quantity)
    return nil
}

func (r TaobaoDrugQuantityUpdateRequest) GetQuantity() int64 {
    return r.quantity
}

