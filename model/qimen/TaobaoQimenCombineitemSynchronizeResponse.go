package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
组合商品接口 APIResponse
taobao.qimen.combineitem.synchronize

ERP调用奇门的接口,将商品信息同步给WMS
*/
type TaobaoQimenCombineitemSynchronizeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenCombineitemSynchronizeResponse `json:"qimen_combineitem_synchronize_response,omitempty"` 
    TaobaoQimenCombineitemSynchronizeResponse
}

/* model for simplify = false
type TaobaoQimenCombineitemSynchronizeResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenCombineitemSynchronizeResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
