package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量获取商品详细信息 APIResponse
taobao.items.seller.list.get

批量获取商品详细信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemsSellerListGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemsSellerListGetResponse `json:"taobao_items_seller_list_get_response,omitempty"`
}

type TaobaoItemsSellerListGetResponse struct {

    // 商品详细信息列表
    Items   []Item `json:"items,omitempty"`

}
