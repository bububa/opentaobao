package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
重命名关键词分组 APIResponse
alibaba.scbp.tag.rename

重命名关键词分组
*/
type AlibabaScbpTagRenameAPIResponse struct {
    model.CommonResponse
    AlibabaScbpTagRenameResponse
}

type AlibabaScbpTagRenameResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_tag_rename_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 重命名分组成功或者失败
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
