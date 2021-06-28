package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除一个推广组 APIResponse
taobao.simba.adgroup.delete

删除一个推广组
*/
type TaobaoSimbaAdgroupDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaAdgroupDeleteResponse `json:"simba_adgroup_delete_response,omitempty"` 
    TaobaoSimbaAdgroupDeleteResponse
}

/* model for simplify = false
type TaobaoSimbaAdgroupDeleteResponse struct {

    // 被删除的推广组
    
    Adgroup  *struct {
        ADGroup  *ADGroup `json:"ad_group,omitempty"`
    } `json:"adgroup,omitempty"`
    

}
*/

type TaobaoSimbaAdgroupDeleteResponse struct {

    // 被删除的推广组
    Adgroup   *ADGroup `json:"adgroup,omitempty"`

}
