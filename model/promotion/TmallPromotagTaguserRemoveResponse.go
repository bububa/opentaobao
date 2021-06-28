package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
给用户移除优惠标签 APIResponse
tmall.promotag.taguser.remove

给用户载体去标
*/
type TmallPromotagTaguserRemoveAPIResponse struct {
    model.CommonResponse
    // Response *TmallPromotagTaguserRemoveResponse `json:"tmall_promotag_taguser_remove_response,omitempty"` 
    TmallPromotagTaguserRemoveResponse
}

/* model for simplify = false
type TmallPromotagTaguserRemoveResponse struct {

    // 打标结果是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TmallPromotagTaguserRemoveResponse struct {

    // 打标结果是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
