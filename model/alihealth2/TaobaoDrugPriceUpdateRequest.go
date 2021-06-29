package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家更新宝贝价格 APIRequest
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

func NewTaobaoDrugPriceUpdateRequest() *TaobaoDrugPriceUpdateRequest{
    return &TaobaoDrugPriceUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDrugPriceUpdateRequest) GetApiMethodName() string {
    return "taobao.drug.price.update"
}

func (r TaobaoDrugPriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDrugPriceUpdateRequest) SetOutStoreId(outStoreId string) error {
    r.outStoreId = outStoreId
    r.Set("out_store_id", outStoreId)
    return nil
}

func (r TaobaoDrugPriceUpdateRequest) GetOutStoreId() string {
    return r.outStoreId
}

func (r *TaobaoDrugPriceUpdateRequest) SetOutItemId(outItemId string) error {
    r.outItemId = outItemId
    r.Set("out_item_id", outItemId)
    return nil
}

func (r TaobaoDrugPriceUpdateRequest) GetOutItemId() string {
    return r.outItemId
}

func (r *TaobaoDrugPriceUpdateRequest) SetPrice(price float64) error {
    r.price = price
    r.Set("price", price)
    return nil
}

func (r TaobaoDrugPriceUpdateRequest) GetPrice() float64 {
    return r.price
}

