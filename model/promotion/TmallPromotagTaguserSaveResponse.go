package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
给用户打上优惠标签 APIResponse
tmall.promotag.taguser.save

给用户载体打标
*/
type TmallPromotagTaguserSaveAPIResponse struct {
    model.CommonResponse
    Response *TmallPromotagTaguserSaveResponse `json:"tmall_promotag_taguser_save_response,omitempty"`
}

type TmallPromotagTaguserSaveResponse struct {

    // 打标结果是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
