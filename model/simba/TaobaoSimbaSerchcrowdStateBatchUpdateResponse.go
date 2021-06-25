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
    Response *TaobaoSimbaSerchcrowdStateBatchUpdateResponse `json:"taobao_simba_serchcrowd_state_batch_update_response,omitempty"`
}

type TaobaoSimbaSerchcrowdStateBatchUpdateResponse struct {

    // 部分失败时返回错误List
    ErrorList   []Json `json:"error_list,omitempty"`

    // result
    Adgrouptargetingtags   []AdgroupTargetingTagDto `json:"adgrouptargetingtags,omitempty"`

}
