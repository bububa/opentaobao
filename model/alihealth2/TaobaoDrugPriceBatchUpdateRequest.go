package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家批量更新宝贝价格 APIRequest
taobao.drug.price.batch.update

商家批量更新宝贝价格
*/
type TaobaoDrugPriceBatchUpdateRequest struct {
    model.Params

    // 外部店铺ID
    outStoreId   string 

    // 商品ID和价格
    outItemIdPriceMap   string 

}

func NewTaobaoDrugPriceBatchUpdateRequest() *TaobaoDrugPriceBatchUpdateRequest{
    return &TaobaoDrugPriceBatchUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDrugPriceBatchUpdateRequest) GetApiMethodName() string {
    return "taobao.drug.price.batch.update"
}

func (r TaobaoDrugPriceBatchUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDrugPriceBatchUpdateRequest) SetOutStoreId(outStoreId string) error {
    r.outStoreId = outStoreId
    r.Set("out_store_id", outStoreId)
    return nil
}

func (r TaobaoDrugPriceBatchUpdateRequest) GetOutStoreId() string {
    return r.outStoreId
}

func (r *TaobaoDrugPriceBatchUpdateRequest) SetOutItemIdPriceMap(outItemIdPriceMap string) error {
    r.outItemIdPriceMap = outItemIdPriceMap
    r.Set("out_item_id_price_map", outItemIdPriceMap)
    return nil
}

func (r TaobaoDrugPriceBatchUpdateRequest) GetOutItemIdPriceMap() string {
    return r.outItemIdPriceMap
}

