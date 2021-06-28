package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取删除的创意ID APIResponse
taobao.simba.creativeids.deleted.get

获取删除的创意ID
*/
type TaobaoSimbaCreativeidsDeletedGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaCreativeidsDeletedGetResponse `json:"simba_creativeids_deleted_get_response,omitempty"` 
    TaobaoSimbaCreativeidsDeletedGetResponse
}

/* model for simplify = false
type TaobaoSimbaCreativeidsDeletedGetResponse struct {

    // 创意ID列表
    
    DeletedCreativeIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"deleted_creative_ids,omitempty"`
    

}
*/

type TaobaoSimbaCreativeidsDeletedGetResponse struct {

    // 创意ID列表
    DeletedCreativeIds   []int64 `json:"deleted_creative_ids,omitempty"`

}
