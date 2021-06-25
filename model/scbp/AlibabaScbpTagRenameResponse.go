package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
重命名关键词分组 APIResponse
alibaba.scbp.tag.rename

重命名关键词分组
*/
type AlibabaScbpTagRenameAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpTagRenameResponse `json:"alibaba_scbp_tag_rename_response,omitempty"`
}

type AlibabaScbpTagRenameResponse struct {

    // 重命名分组成功或者失败
    Result   bool `json:"result,omitempty"`

}
