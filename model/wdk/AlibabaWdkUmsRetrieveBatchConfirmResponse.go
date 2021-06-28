package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量消息确认 APIResponse
alibaba.wdk.ums.retrieve.batch.confirm

批量消息确认
*/
type AlibabaWdkUmsRetrieveBatchConfirmAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_ums_retrieve_batch_confirm_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *UtmsResult `json:"result,omitempty" xml:"