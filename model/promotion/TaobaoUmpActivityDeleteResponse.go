package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除营销活动 APIResponse
taobao.ump.activity.delete

删除营销活动。对应的活动详情等将会被全部删除。
*/
type TaobaoUmpActivityDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoUmpActivityDeleteResponse `json:"ump_activity_delete_response,omitempty"` 
    TaobaoUmpActivityDeleteResponse
}

/* model for simplify = false
type TaobaoUmpActivityDeleteResponse struct {

    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoUmpActivityDeleteResponse struct {

    // 调用是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
