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
    Response *TaobaoSimbaSearchcrowdBatchAddResponse `json:"taobao_simba_searchcrowd_batch_add_response,omitempty"`
}

type TaobaoSimbaSearchcrowdBatchAddResponse struct {

    // 定向信息
    Adgrouptargetingtags   []AdgroupTargetingTagDto `json:"adgrouptargetingtags,omitempty"`

}
