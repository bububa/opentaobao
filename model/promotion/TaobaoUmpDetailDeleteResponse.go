package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除活动详情 APIResponse
taobao.ump.detail.delete

删除活动详情
*/
type TaobaoUmpDetailDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoUmpDetailDeleteResponse `json:"ump_detail_delete_response,omitempty"` 
    TaobaoUmpDetailDeleteResponse
}

/* model for simplify = false
type TaobaoUmpDetailDeleteResponse struct {

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoUmpDetailDeleteResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
