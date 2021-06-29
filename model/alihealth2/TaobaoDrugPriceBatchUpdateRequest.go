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
    _outStoreId   string
    // 商品ID和价格
    _outItemIdPriceMap   string
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
func (r *TaobaoDrugPriceBatchUpdateRequest) SetOutStoreId(_outStoreId string) error {
    r._outStoreId = _outStoreId
    r.Set("out_store_id", _outStoreId)
    return nil
}

// OutStoreId Getter
func (r TaobaoDrugPriceBatchUpdateRequest) GetOutStoreId() string {
    return r._outStoreId
}
// OutItemIdPriceMap Setter
// 商品ID和价格
func (r *TaobaoDrugPriceBatchUpdateRequest) SetOutItemIdPriceMap(_outItemIdPriceMap string) error {
    r._outItemIdPriceMap = _outItemIdPriceMap
    r.Set("out_item_id_price_map", _outItemIdPriceMap)
    return nil
}

// OutItemIdPriceMap Getter
func (r TaobaoDrugPriceBatchUpdateRequest) GetOutItemIdPriceMap() string {
    return r._outItemIdPriceMap
}
