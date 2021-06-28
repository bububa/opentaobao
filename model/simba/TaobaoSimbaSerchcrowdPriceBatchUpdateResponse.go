package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
单品推广搜索人群修改溢价 APIResponse
taobao.simba.serchcrowd.price.batch.update

单品推广搜索人群修改溢价, 不支持跨推广单元修改
*/
type TaobaoSimbaSerchcrowdPriceBatchUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaSerchcrowdPriceBatchUpdateResponse `json:"simba_serchcrowd_price_batch_update_response,omitempty"` 
    TaobaoSimbaSerchcrowdPriceBatchUpdateResponse
}

/* model for simplify = false
type TaobaoSimbaSerchcrowdPriceBatchUpdateResponse struct {

    // result
    
    Adgrouptargetingtags  struct {
        AdgroupTargetingTagDto  []AdgroupTargetingTagDto `json:"adgroup_targeting_tag_dto,omitempty"`
    } `json:"adgrouptargetingtags,omitempty"`
    

}
*/

type TaobaoSimbaSerchcrowdPriceBatchUpdateResponse struct {

    // result
    Adgrouptargetingtags   []AdgroupTargetingTagDto `json:"adgrouptargetingtags,omitempty"`

}
