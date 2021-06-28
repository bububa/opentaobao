package shop

import (
    "github.com/bububa/opentaobao/model"
)

/* 
卖家店铺基础信息查询 APIResponse
taobao.shop.seller.get

获取卖家店铺的基本信息
*/
type TaobaoShopSellerGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoShopSellerGetResponse `json:"shop_seller_get_response,omitempty"` 
    TaobaoShopSellerGetResponse
}

/* model for simplify = false
type TaobaoShopSellerGetResponse struct {

    // 店铺信息
    
    Shop  *struct {
        Shop  *Shop `json:"shop,omitempty"`
    } `json:"shop,omitempty"`
    

}
*/

type TaobaoShopSellerGetResponse struct {

    // 店铺信息
    Shop   *Shop `json:"shop,omitempty"`

}
