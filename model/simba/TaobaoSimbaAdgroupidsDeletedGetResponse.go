package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取删除的推广组ID APIResponse
taobao.simba.adgroupids.deleted.get

获取删除的推广组ID
*/
type TaobaoSimbaAdgroupidsDeletedGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaAdgroupidsDeletedGetResponse `json:"simba_adgroupids_deleted_get_response,omitempty"` 
    TaobaoSimbaAdgroupidsDeletedGetResponse
}

/* model for simplify = false
type TaobaoSimbaAdgroupidsDeletedGetResponse struct {

    // 推广组ID列表
    
    DeletedAdgroupIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"deleted_adgroup_ids,omitempty"`
    

}
*/

type TaobaoSimbaAdgroupidsDeletedGetResponse struct {

    // 推广组ID列表
    DeletedAdgroupIds   []int64 `json:"deleted_adgroup_ids,omitempty"`

}
