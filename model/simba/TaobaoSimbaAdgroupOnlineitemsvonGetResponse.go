package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取用户上架在线销售的全部宝贝 APIResponse
taobao.simba.adgroup.onlineitemsvon.get

获取用户上架在线销售的全部宝贝
*/
type TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaAdgroupOnlineitemsvonGetResponse `json:"simba_adgroup_onlineitemsvon_get_response,omitempty"` 
    TaobaoSimbaAdgroupOnlineitemsvonGetResponse
}

/* model for simplify = false
type TaobaoSimbaAdgroupOnlineitemsvonGetResponse struct {

    // 带分页的淘宝商品
    
    PageItem  *struct {
        SubwayItemPartition  *SubwayItemPartition `json:"subway_item_partition,omitempty"`
    } `json:"page_item,omitempty"`
    

}
*/

type TaobaoSimbaAdgroupOnlineitemsvonGetResponse struct {

    // 带分页的淘宝商品
    PageItem   *SubwayItemPartition `json:"page_item,omitempty"`

}
