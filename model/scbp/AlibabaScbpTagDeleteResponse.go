package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除关键词分组 APIResponse
alibaba.scbp.tag.delete

删除关键词分组
*/
type AlibabaScbpTagDeleteAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpTagDeleteResponse `json:"alibaba_scbp_tag_delete_response,omitempty"`
}

type AlibabaScbpTagDeleteResponse struct {

    // 删除关键词分组成功
    Result   bool `json:"result,omitempty"`

}
