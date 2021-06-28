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
    // Response *TaobaoSimbaSerchcrowdBatchDeleteResponse `json:"simba_serchcrowd_batch_delete_response,omitempty"` 
    TaobaoSimbaSerchcrowdBatchDeleteResponse
}

/* model for simplify = false
type TaobaoSimbaSerchcrowdBatchDeleteResponse struct {

    // result
    
    DeleteList  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"delete_list,omitempty"`
    

    // errorDTOList
    
    ErrorDTOList  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"error_d_t_o_list,omitempty"`
    

    // key
    
    Key   string `json:"key,omitempty"`
    

}
*/

type TaobaoSimbaSerchcrowdBatchDeleteResponse struct {

    // result
    DeleteList   []string `json:"delete_list,omitempty"`

    // errorDTOList
    ErrorDTOList   []string `json:"error_d_t_o_list,omitempty"`

    // key
    Key   string `json:"key,omitempty"`

}
