package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取单个商品详细信息 APIResponse
taobao.item.seller.get

获取单个商品的全部信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemSellerGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoItemSellerGetResponse `json:"item_seller_get_response,omitempty"` 
    TaobaoItemSellerGetResponse
}

/* model for simplify = false
type TaobaoItemSellerGetResponse struct {

    // 商品详细信息
    
    Item  *struct {
        Item  *Item `json:"item,omitempty"`
    } `json:"item,omitempty"`
    

}
*/

type TaobaoItemSellerGetResponse struct {

    // 商品详细信息
    Item   *Item `json:"item,omitempty"`

}
