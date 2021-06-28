package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品同步接口 (批量) APIResponse
taobao.qimen.items.synchronize

ERP调用奇门的接口,批量同步商品信息给WMS
*/
type TaobaoQimenItemsSynchronizeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenItemsSynchronizeResponse `json:"qimen_items_synchronize_response,omitempty"` 
    TaobaoQimenItemsSynchronizeResponse
}

/* model for simplify = false
type TaobaoQimenItemsSynchronizeResponse struct {

    // 
    
    Response  *struct {
        ItemsSynchronizeResponse  *ItemsSynchronizeResponse `json:"items_synchronize_response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenItemsSynchronizeResponse struct {

    // 
    Response   *ItemsSynchronizeResponse `json:"response,omitempty"`

}
