package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新/增加sku信息 APIRequest
alibaba.idle.rent.item.sku.update

更新/增加sku信息
*/
type AlibabaIdleRentItemSkuUpdateRequest struct {
    model.Params

    // sku信息，更新后skuId保持不变
    sku   *ItemSkuDto 

}

func NewAlibabaIdleRentItemSkuUpdateRequest() *AlibabaIdleRentItemSkuUpdateRequest{
    return &AlibabaIdleRentItemSkuUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRentItemSkuUpdateRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.item.sku.update"
}

func (r AlibabaIdleRentItemSkuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRentItemSkuUpdateRequest) SetSku(sku *ItemSkuDto) error {
    r.sku = sku
    r.Set("sku", sku)
    return nil
}

func (r AlibabaIdleRentItemSkuUpdateRequest) GetSku() *ItemSkuDto {
    return r.sku
}

