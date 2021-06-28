package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广组的信息 APIResponse
taobao.simba.adgroup.update

更新一个推广组的信息，可以设置默认出价、是否上线、非搜索出价、非搜索是否使用默认出价
*/
type TaobaoSimbaAdgroupUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaAdgroupUpdateResponse `json:"simba_adgroup_update_response,omitempty"` 
    TaobaoSimbaAdgroupUpdateResponse
}

/* model for simplify = false
type TaobaoSimbaAdgroupUpdateResponse struct {

    // 被修改的推广组
    
    Adgroup  *struct {
        ADGroup  *ADGroup `json:"ad_group,omitempty"`
    } `json:"adgroup,omitempty"`
    

}
*/

type TaobaoSimbaAdgroupUpdateResponse struct {

    // 被修改的推广组
    Adgroup   *ADGroup `json:"adgroup,omitempty"`

}
