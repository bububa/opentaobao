package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推广单元增加搜索人群 APIResponse
taobao.simba.searchcrowd.batch.add

推广单元新增搜索人群
*/
type TaobaoSimbaSearchcrowdBatchAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaSearchcrowdBatchAddResponse `json:"simba_searchcrowd_batch_add_response,omitempty"` 
    TaobaoSimbaSearchcrowdBatchAddResponse
}

/* model for simplify = false
type TaobaoSimbaSearchcrowdBatchAddResponse struct {

    // 定向信息
    
    Adgrouptargetingtags  struct {
        AdgroupTargetingTagDto  []AdgroupTargetingTagDto `json:"adgroup_targeting_tag_dto,omitempty"`
    } `json:"adgrouptargetingtags,omitempty"`
    

}
*/

type TaobaoSimbaSearchcrowdBatchAddResponse struct {

    // 定向信息
    Adgrouptargetingtags   []AdgroupTargetingTagDto `json:"adgrouptargetingtags,omitempty"`

}
