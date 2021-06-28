package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建关键词分组 APIResponse
alibaba.scbp.tag.add

创建关键词分组
*/
type AlibabaScbpTagAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_tag_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分组名称
    
    TagName   string `json:"tag_name,omitempty" xml:"