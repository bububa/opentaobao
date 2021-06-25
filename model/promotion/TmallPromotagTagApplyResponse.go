package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
优惠标签申请 APIResponse
tmall.promotag.tag.apply

创建优惠标签
*/
type TmallPromotagTagApplyAPIResponse struct {
    model.CommonResponse
    Response *TmallPromotagTagApplyResponse `json:"tmall_promotag_tag_apply_response,omitempty"`
}

type TmallPromotagTagApplyResponse struct {

    // 是否设置成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 优惠标签ID
    TagId   int64 `json:"tag_id,omitempty"`

}
