package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建关键词分组 APIResponse
alibaba.scbp.tag.add

创建关键词分组
*/
type AlibabaScbpTagAddAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpTagAddResponse `json:"alibaba_scbp_tag_add_response,omitempty"` 
    AlibabaScbpTagAddResponse
}

/* model for simplify = false
type AlibabaScbpTagAddResponse struct {

    // 分组名称
    
    TagName   string `json:"tag_name,omitempty"`
    

}
*/

type AlibabaScbpTagAddResponse struct {

    // 分组名称
    TagName   string `json:"tag_name,omitempty"`

}
