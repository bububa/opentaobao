package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家更新宝贝价格 API请求
taobao.drug.price.update

商家更新价格
*/
type TaobaoDrugPriceUpdateRequest struct {
    model.Params
    // 对应的外部店铺ID
    outStoreId   string
    // 对应的外部商品编码
    outItemId   string
    // 商品价格
    price   float64
}

// 初始化TaobaoDrugPriceUpdateRequest对象
func NewTaobaoDrugPriceUpdateRequest() *TaobaoDrugPriceUpdateRequest{
    return &TaobaoDrugPriceUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDrugPriceUpdateRequest) GetApiMethodName() string {
    return "taobao.drug.price.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDrugPriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutStoreId Setter
// 对应的外部店铺ID
func (r *TaobaoDrugPriceUpdateRequest) SetOutStoreId(outStoreId string) error {
    r.outStoreId = outStoreId
    r.Set("out_store_id", outStoreId)
    return nil
}

// OutStoreId Getter
func (r TaobaoDrugPriceUpdateRequest) GetOutStoreId() string {
    return r.outStoreId
}
// OutItemId Setter
// 对应的外部商品编码
func (r *TaobaoDrugPriceUpdateRequest) SetOutItemId(outItemId string) error {
    r.outItemId = outItemId
    r.Set("out_item_id", outItemId)
    return nil
}

// OutItemId Getter
func (r TaobaoDrugPriceUpdateRequest) GetOutItemId() string {
    return r.outItemId
}
// Price Setter
// 商品价格
func (r *TaobaoDrugPriceUpdateRequest) SetPrice(price float64) error {
    r.price = price
    r.Set("price", price)
    return nil
}

// Price Getter
func (r TaobaoDrugPriceUpdateRequest) GetPrice() float64 {
    return r.price
}
