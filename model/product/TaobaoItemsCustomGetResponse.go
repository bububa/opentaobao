package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据外部ID取商品 APIResponse
taobao.items.custom.get

跟据卖家设定的商品外部id获取商品，只能获取授权卖家的商品
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemsCustomGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoItemsCustomGetResponse `json:"items_custom_get_response,omitempty"` 
    TaobaoItemsCustomGetResponse
}

/* model for simplify = false
type TaobaoItemsCustomGetResponse struct {

    // 商品列表，具体返回字段以fields决定
    
    Items  struct {
        Item  []Item `json:"item,omitempty"`
    } `json:"items,omitempty"`
    

}
*/

type TaobaoItemsCustomGetResponse struct {

    // 商品列表，具体返回字段以fields决定
    Items   []Item `json:"items,omitempty"`

}
