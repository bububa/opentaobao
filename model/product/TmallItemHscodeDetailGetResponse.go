package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通过hscode获取计量单位 APIResponse
tmall.item.hscode.detail.get

通过hscode获取计量单位和销售单位
*/
type TmallItemHscodeDetailGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemHscodeDetailGetResponse `json:"tmall_item_hscode_detail_get_response,omitempty"` 
    TmallItemHscodeDetailGetResponse
}

/* model for simplify = false
type TmallItemHscodeDetailGetResponse struct {

    // 返回的计量单位和销售单位
    
    Results  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type TmallItemHscodeDetailGetResponse struct {

    // 返回的计量单位和销售单位
    Results   []string `json:"results,omitempty"`

}
