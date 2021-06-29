package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量同步库存接口 APIRequest
taobao.drug.quantity.batch.update

商家通过top接口可以批量修改商品库存
*/
type TaobaoDrugQuantityBatchUpdateRequest struct {
    model.Params

    // 外部店铺ID
    outStoreId   string 

    // 商品ID和库存
    outItemIdQuantityMap   string 

}

func NewTaobaoDrugQuantityBatchUpdateRequest() *TaobaoDrugQuantityBatchUpdateRequest{
    return &TaobaoDrugQuantityBatchUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDrugQuantityBatchUpdateRequest) GetApiMethodName() string {
    return "taobao.drug.quantity.batch.update"
}

func (r TaobaoDrugQuantityBatchUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDrugQuantityBatchUpdateRequest) SetOutStoreId(outStoreId string) error {
    r.outStoreId = outStoreId
    r.Set("out_store_id", outStoreId)
    return nil
}

func (r TaobaoDrugQuantityBatchUpdateRequest) GetOutStoreId() string {
    return r.outStoreId
}

func (r *TaobaoDrugQuantityBatchUpdateRequest) SetOutItemIdQuantityMap(outItemIdQuantityMap string) error {
    r.outItemIdQuantityMap = outItemIdQuantityMap
    r.Set("out_item_id_quantity_map", outItemIdQuantityMap)
    return nil
}

func (r TaobaoDrugQuantityBatchUpdateRequest) GetOutItemIdQuantityMap() string {
    return r.outItemIdQuantityMap
}

