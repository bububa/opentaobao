package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品是否推广 APIResponse
taobao.simba.adgroups.item.exist

判断在一个推广计划中是否已经推广了一个商品
*/
type TaobaoSimbaAdgroupsItemExistAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaAdgroupsItemExistResponse `json:"simba_adgroups_item_exist_response,omitempty"` 
    TaobaoSimbaAdgroupsItemExistResponse
}

/* model for simplify = false
type TaobaoSimbaAdgroupsItemExistResponse struct {

    // true表示已经被推广，false表示没有被推广
    
    Exist   bool `json:"exist,omitempty"`
    

}
*/

type TaobaoSimbaAdgroupsItemExistResponse struct {

    // true表示已经被推广，false表示没有被推广
    Exist   bool `json:"exist,omitempty"`

}
