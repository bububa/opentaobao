package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商户门店创建接口 APIResponse
taobao.place.store.create

用于商家创建线下门店
*/
type TaobaoPlaceStoreCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPlaceStoreCreateResponse `json:"place_store_create_response,omitempty"` 
    TaobaoPlaceStoreCreateResponse
}

/* model for simplify = false
type TaobaoPlaceStoreCreateResponse struct {

    // result
    
    StoreId   int64 `json:"store_id,omitempty"`
    

}
*/

type TaobaoPlaceStoreCreateResponse struct {

    // result
    StoreId   int64 `json:"store_id,omitempty"`

}
