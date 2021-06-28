package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询标签接口 APIResponse
tmall.promotag.tag.find

查询用户创建的所有标签
*/
type TmallPromotagTagFindAPIResponse struct {
    model.CommonResponse
    TmallPromotagTagFindResponse
}

type TmallPromotagTagFindResponse struct {
    XMLName xml.Name `xml:"tmall_promotag_tag_find_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询结果类型
    
    QueryResult   *PromotionTagQuery `json:"query_result,omitempty" xml:"query_result,omitempty"`

    
}
