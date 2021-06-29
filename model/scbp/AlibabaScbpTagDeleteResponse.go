package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除关键词分组 APIResponse
alibaba.scbp.tag.delete

删除关键词分组
*/
type AlibabaScbpTagDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaScbpTagDeleteResponse
}

type AlibabaScbpTagDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_tag_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 删除关键词分组成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
