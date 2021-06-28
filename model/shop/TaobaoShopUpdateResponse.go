package shop

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新店铺基本信息 APIResponse
taobao.shop.update

目前只支持标题、公告和描述的更新
*/
type TaobaoShopUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoShopUpdateResponse `json:"shop_update_response,omitempty"` 
    TaobaoShopUpdateResponse
}

/* model for simplify = false
type TaobaoShopUpdateResponse struct {

    // 店铺信息
    
    Shop  *struct {
        Shop  *Shop `json:"shop,omitempty"`
    } `json:"shop,omitempty"`
    

}
*/

type TaobaoShopUpdateResponse struct {

    // 店铺信息
    Shop   *Shop `json:"shop,omitempty"`

}
