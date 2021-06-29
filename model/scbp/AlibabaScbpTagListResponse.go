package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询所有分组 APIResponse
alibaba.scbp.tag.list

查询所有分组
*/
type AlibabaScbpTagListAPIResponse struct {
    model.CommonResponse
    AlibabaScbpTagListResponse
}

type AlibabaScbpTagListResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_tag_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 所有分组
    
    TagList   []TagGroup `json:"tag_list,omitempty" xml:"tag_list>tag_group,omitempty"`
    
    
}
