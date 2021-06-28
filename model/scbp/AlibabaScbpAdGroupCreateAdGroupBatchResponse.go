package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建推广单元 APIResponse
alibaba.scbp.ad.group.create.ad.group.batch

创建推广单元
*/
type AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdGroupCreateAdGroupBatchResponse
}

type AlibabaScbpAdGroupCreateAdGroupBatchResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_group_create_ad_group_batch_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
