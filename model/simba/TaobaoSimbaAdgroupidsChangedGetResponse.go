package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取修改的推广组ID APIResponse
taobao.simba.adgroupids.changed.get

获取修改的推广组ID
*/
type TaobaoSimbaAdgroupidsChangedGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaAdgroupidsChangedGetResponse `json:"simba_adgroupids_changed_get_response,omitempty"` 
    TaobaoSimbaAdgroupidsChangedGetResponse
}

/* model for simplify = false
type TaobaoSimbaAdgroupidsChangedGetResponse struct {

    // 推广组ID列表
    
    ChangedAdgroupids  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"changed_adgroupids,omitempty"`
    

}
*/

type TaobaoSimbaAdgroupidsChangedGetResponse struct {

    // 推广组ID列表
    ChangedAdgroupids   []int64 `json:"changed_adgroupids,omitempty"`

}
