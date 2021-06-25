package shop

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取前台展示的店铺内卖家自定义商品类目 APIResponse
taobao.sellercats.list.get

此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家）
*/
type TaobaoSellercatsListGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSellercatsListGetResponse `json:"taobao_sellercats_list_get_response,omitempty"`
}

type TaobaoSellercatsListGetResponse struct {

    // 卖家自定义类目
    SellerCats   []SellerCat `json:"seller_cats,omitempty"`

}
