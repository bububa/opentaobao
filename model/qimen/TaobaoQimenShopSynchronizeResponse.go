package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
店铺同步接口 APIResponse
taobao.qimen.shop.synchronize

店铺同步接口描述
*/
type TaobaoQimenShopSynchronizeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenShopSynchronizeResponse `json:"qimen_shop_synchronize_response,omitempty"` 
    TaobaoQimenShopSynchronizeResponse
}

/* model for simplify = false
type TaobaoQimenShopSynchronizeResponse struct {

    // Response
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenShopSynchronizeResponse struct {

    // Response
    Response   *Response `json:"response,omitempty"`

}
