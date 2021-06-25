package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
单品搜索人群批量取消溢价 APIResponse
taobao.simba.serchcrowd.batch.delete

删除单品搜索人群溢价功能
*/
type TaobaoSimbaSerchcrowdBatchDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaSerchcrowdBatchDeleteResponse `json:"taobao_simba_serchcrowd_batch_delete_response,omitempty"`
}

type TaobaoSimbaSerchcrowdBatchDeleteResponse struct {

    // result
    DeleteList   []Json `json:"delete_list,omitempty"`

    // errorDTOList
    ErrorDTOList   []Json `json:"error_d_t_o_list,omitempty"`

    // key
    Key   string `json:"key,omitempty"`

}
