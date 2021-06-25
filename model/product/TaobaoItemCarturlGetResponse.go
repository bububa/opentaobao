package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
加购URL获取 APIResponse
taobao.item.carturl.get

获取加购URL，支持添加商品到购物车
*/
type TaobaoItemCarturlGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemCarturlGetResponse `json:"taobao_item_carturl_get_response,omitempty"`
}

type TaobaoItemCarturlGetResponse struct {

    // 加购的URL地址
    Result   string `json:"result,omitempty"`

}
