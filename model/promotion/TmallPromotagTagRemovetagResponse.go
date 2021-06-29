package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除标签定义 API返回值 
tmall.promotag.tag.removetag

用于删除标签定义，但是要确保目前该标签没有人群在使用。
*/
type TmallPromotagTagRemovetagAPIResponse struct {
    model.CommonResponse
    TmallPromotagTagRemovetagResponse
}

// 删除标签定义 成功返回结果
type TmallPromotagTagRemovetagResponse struct {
    XMLName xml.Name `xml:"tmall_promotag_tag_removetag_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 删除操作是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
