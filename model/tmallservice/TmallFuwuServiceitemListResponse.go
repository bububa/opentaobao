package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取服务商品扩展信息 APIResponse
tmall.fuwu.serviceitem.list

获取服务商品扩展信息
*/
type TmallFuwuServiceitemListAPIResponse struct {
    model.CommonResponse
    // Response *TmallFuwuServiceitemListResponse `json:"tmall_fuwu_serviceitem_list_response,omitempty"` 
    TmallFuwuServiceitemListResponse
}

/* model for simplify = false
type TmallFuwuServiceitemListResponse struct {

    // result
    
    Result  *struct {
        TmallFuwuServiceitemListResult  *TmallFuwuServiceitemListResult `json:"tmall_fuwu_serviceitem_list_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallFuwuServiceitemListResponse struct {

    // result
    Result   *TmallFuwuServiceitemListResult `json:"result,omitempty"`

}
