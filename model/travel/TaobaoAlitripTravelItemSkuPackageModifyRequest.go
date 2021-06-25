package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】套餐级别日历价格库存增删操作 APIRequest
taobao.alitrip.travel.item.sku.package.modify

【API3.0】套餐级别日历价格库存增删操作
*/
type TaobaoAlitripTravelItemSkuPackageModifyRequest struct {
    model.Params

    // 商品id。itemId和outProductId至少填写一个
    itemId   int64 

    // 商品 外部商家编码。itemId和outProductId至少填写一个
    outProductId   string 

    // 商品日历价格库存套餐
    skus   []ItemSkuInfo 

}

func NewTaobaoAlitripTravelItemSkuPackageModifyRequest() *TaobaoAlitripTravelItemSkuPackageModifyRequest{
    return &TaobaoAlitripTravelItemSkuPackageModifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelItemSkuPackageModifyRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.item.sku.package.modify"
}

func (r TaobaoAlitripTravelItemSkuPackageModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelItemSkuPackageModifyRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoAlitripTravelItemSkuPackageModifyRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoAlitripTravelItemSkuPackageModifyRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

func (r TaobaoAlitripTravelItemSkuPackageModifyRequest) GetOutProductId() string {
    return r.outProductId
}

func (r *TaobaoAlitripTravelItemSkuPackageModifyRequest) SetSkus(skus []ItemSkuInfo) error {
    r.skus = skus
    r.Set("skus", skus)
    return nil
}

func (r TaobaoAlitripTravelItemSkuPackageModifyRequest) GetSkus() []ItemSkuInfo {
    return r.skus
}

