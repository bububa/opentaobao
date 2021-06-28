package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询所有分组 APIResponse
alibaba.scbp.tag.list

查询所有分组
*/
type AlibabaScbpTagListAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpTagListResponse `json:"alibaba_scbp_tag_list_response,omitempty"` 
    AlibabaScbpTagListResponse
}

/* model for simplify = false
type AlibabaScbpTagListResponse struct {

    // 所有分组
    
    TagList  struct {
        TagGroup  []TagGroup `json:"tag_group,omitempty"`
    } `json:"tag_list,omitempty"`
    

}
*/

type AlibabaScbpTagListResponse struct {

    // 所有分组
    TagList   []TagGroup `json:"tag_list,omitempty"`

}
