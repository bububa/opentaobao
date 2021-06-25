package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新商品价格 APIResponse
taobao.item.price.update

更新商品价格
*/
type TaobaoItemPriceUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemPriceUpdateResponse `json:"taobao_item_price_update_response,omitempty"`
}

type TaobaoItemPriceUpdateResponse struct {

    // 商品结构里的num_iid，modified
    Item   *Item `json:"item,omitempty"`

}
