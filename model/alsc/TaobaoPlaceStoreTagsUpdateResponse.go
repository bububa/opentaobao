package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店打标去标 APIResponse
taobao.place.store.tags.update

门店打标去标
*/
type TaobaoPlaceStoreTagsUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPlaceStoreTagsUpdateResponse `json:"place_store_tags_update_response,omitempty"` 
    TaobaoPlaceStoreTagsUpdateResponse
}

/* model for simplify = false
type TaobaoPlaceStoreTagsUpdateResponse struct {

    // 返回结果：true成功；false失败
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type TaobaoPlaceStoreTagsUpdateResponse struct {

    // 返回结果：true成功；false失败
    Result   bool `json:"result,omitempty"`

}
