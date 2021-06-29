package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建推广单元 API返回值 
alibaba.scbp.ad.group.create.ad.group.batch

创建推广单元
*/
type AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdGroupCreateAdGroupBatchResponse
}

// 创建推广单元 成功返回结果
type AlibabaScbpAdGroupCreateAdGroupBatchResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_group_create_ad_group_batch_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
