package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线下门店删除 APIResponse
taobao.place.store.delete

用于商家删除线下门店
*/
type TaobaoPlaceStoreDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPlaceStoreDeleteResponse `json:"place_store_delete_response,omitempty"` 
    TaobaoPlaceStoreDeleteResponse
}

/* model for simplify = false
type TaobaoPlaceStoreDeleteResponse struct {

    // 门店删除结果
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type TaobaoPlaceStoreDeleteResponse struct {

    // 门店删除结果
    Result   bool `json:"result,omitempty"`

}
