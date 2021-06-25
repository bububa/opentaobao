package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新商品SKU的价格 APIResponse
taobao.item.sku.price.update

更新商品SKU的价格
*/
type TaobaoItemSkuPriceUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemSkuPriceUpdateResponse `json:"taobao_item_sku_price_update_response,omitempty"`
}

type TaobaoItemSkuPriceUpdateResponse struct {

    // 商品SKU信息（只包含num_iid和modified）
    Sku   *Sku `json:"sku,omitempty"`

}
