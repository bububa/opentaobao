package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家批量更新宝贝价格 API请求
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

// 初始化TaobaoDrugPriceBatchUpdateRequest对象
func NewTaobaoDrugPriceBatchUpdateRequest() *TaobaoDrugPriceBatchUpdateRequest{
    return &TaobaoDrugPriceBatchUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDrugPriceBatchUpdateRequest) GetApiMethodName() string {
    return "taobao.drug.price.batch.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDrugPriceBatchUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutStoreId Setter
// 外部店铺ID
func (r *TaobaoDrugPriceBatchUpdateRequest) SetOutStoreId(outStoreId string) error {
    r.outStoreId = outStoreId
    r.Set("out_store_id", outStoreId)
    return nil
}

// OutStoreId Getter
func (r TaobaoDrugPriceBatchUpdateRequest) GetOutStoreId() string {
    return r.outStoreId
}
// OutItemIdPriceMap Setter
// 商品ID和价格
func (r *TaobaoDrugPriceBatchUpdateRequest) SetOutItemIdPriceMap(outItemIdPriceMap string) error {
    r.outItemIdPriceMap = outItemIdPriceMap
    r.Set("out_item_id_price_map", outItemIdPriceMap)
    return nil
}

// OutItemIdPriceMap Getter
func (r TaobaoDrugPriceBatchUpdateRequest) GetOutItemIdPriceMap() string {
    return r.outItemIdPriceMap
}
