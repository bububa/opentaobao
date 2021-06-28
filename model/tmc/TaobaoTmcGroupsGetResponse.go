package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取自定义用户分组列表 APIResponse
taobao.tmc.groups.get

获取自定义用户分组列表
*/
type TaobaoTmcGroupsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTmcGroupsGetResponse `json:"tmc_groups_get_response,omitempty"` 
    TaobaoTmcGroupsGetResponse
}

/* model for simplify = false
type TaobaoTmcGroupsGetResponse struct {

    // dasdasd
    
    Groups  struct {
        TmcGroup  []TmcGroup `json:"tmc_group,omitempty"`
    } `json:"groups,omitempty"`
    

    // 分组总数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

}
*/

type TaobaoTmcGroupsGetResponse struct {

    // dasdasd
    Groups   []TmcGroup `json:"groups,omitempty"`

    // 分组总数
    TotalResults   int64 `json:"total_results,omitempty"`

}
