package shop

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新卖家自定义类目 APIResponse
taobao.sellercats.list.update

此API更新卖家店铺内自定义类目 <br/>注：因为缓存的关系，添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布
*/
type TaobaoSellercatsListUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSellercatsListUpdateResponse `json:"sellercats_list_update_response,omitempty"` 
    TaobaoSellercatsListUpdateResponse
}

/* model for simplify = false
type TaobaoSellercatsListUpdateResponse struct {

    // 返回sellercat数据结构中的：cid,modified
    
    SellerCat  *struct {
        SellerCat  *SellerCat `json:"seller_cat,omitempty"`
    } `json:"seller_cat,omitempty"`
    

}
*/

type TaobaoSellercatsListUpdateResponse struct {

    // 返回sellercat数据结构中的：cid,modified
    SellerCat   *SellerCat `json:"seller_cat,omitempty"`

}
