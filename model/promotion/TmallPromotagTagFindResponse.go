package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询标签接口 APIResponse
tmall.promotag.tag.find

查询用户创建的所有标签
*/
type TmallPromotagTagFindAPIResponse struct {
    model.CommonResponse
    Response *TmallPromotagTagFindResponse `json:"tmall_promotag_tag_find_response,omitempty"`
}

type TmallPromotagTagFindResponse struct {

    // 查询结果类型
    QueryResult   *PromotionTagQuery `json:"query_result,omitempty"`

}
