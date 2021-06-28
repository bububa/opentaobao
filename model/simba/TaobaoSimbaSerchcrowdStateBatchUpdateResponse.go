package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
单品搜索人群修改状态 APIResponse
taobao.simba.serchcrowd.state.batch.update

暂停或启用单品推广搜索人群溢价
*/
type TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaSerchcrowdStateBatchUpdateResponse `json:"simba_serchcrowd_state_batch_update_response,omitempty"` 
    TaobaoSimbaSerchcrowdStateBatchUpdateResponse
}

/* model for simplify = false
type TaobaoSimbaSerchcrowdStateBatchUpdateResponse struct {

    // 部分失败时返回错误List
    
    ErrorList  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"error_list,omitempty"`
    

    // result
    
    Adgrouptargetingtags  struct {
        AdgroupTargetingTagDto  []AdgroupTargetingTagDto `json:"adgroup_targeting_tag_dto,omitempty"`
    } `json:"adgrouptargetingtags,omitempty"`
    

}
*/

type TaobaoSimbaSerchcrowdStateBatchUpdateResponse struct {

    // 部分失败时返回错误List
    ErrorList   []string `json:"error_list,omitempty"`

    // result
    Adgrouptargetingtags   []AdgroupTargetingTagDto `json:"adgrouptargetingtags,omitempty"`

}
