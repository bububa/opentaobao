package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分页获取修改的推广组ID和修改时间 APIResponse
taobao.simba.adgroups.changed.get

分页获取修改的推广组ID和修改时间
*/
type TaobaoSimbaAdgroupsChangedGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaAdgroupsChangedGetResponse `json:"simba_adgroups_changed_get_response,omitempty"` 
    TaobaoSimbaAdgroupsChangedGetResponse
}

/* model for simplify = false
type TaobaoSimbaAdgroupsChangedGetResponse struct {

    // 推广组分页对象
    
    Adgroups  *struct {
        ADGroupPage  *ADGroupPage `json:"ad_group_page,omitempty"`
    } `json:"adgroups,omitempty"`
    

}
*/

type TaobaoSimbaAdgroupsChangedGetResponse struct {

    // 推广组分页对象
    Adgroups   *ADGroupPage `json:"adgroups,omitempty"`

}
