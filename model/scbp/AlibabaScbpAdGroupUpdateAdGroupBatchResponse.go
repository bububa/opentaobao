package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改推广单元 APIResponse
alibaba.scbp.ad.group.update.ad.group.batch

修改推广单元
*/
type AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_group_update_ad_group_batch_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   int64 `json:"result,omitempty" xml:"