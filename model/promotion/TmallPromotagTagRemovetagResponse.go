package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除标签定义 APIResponse
tmall.promotag.tag.removetag

用于删除标签定义，但是要确保目前该标签没有人群在使用。
*/
type TmallPromotagTagRemovetagAPIResponse struct {
    model.CommonResponse
    // Response *TmallPromotagTagRemovetagResponse `json:"tmall_promotag_tag_removetag_response,omitempty"` 
    TmallPromotagTagRemovetagResponse
}

/* model for simplify = false
type TmallPromotagTagRemovetagResponse struct {

    // 删除操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TmallPromotagTagRemovetagResponse struct {

    // 删除操作是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
