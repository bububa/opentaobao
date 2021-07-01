package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改推广单元 API返回值 
alibaba.scbp.ad.group.update.ad.group.batch

修改推广单元
*/
type AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponseModel
}

// 修改推广单元 成功返回结果
type AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_group_update_ad_group_batch_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`
}
