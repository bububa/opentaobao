package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量消息确认 APIResponse
alibaba.wdk.ums.retrieve.batch.confirm

批量消息确认
*/
type AlibabaWdkUmsRetrieveBatchConfirmAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkUmsRetrieveBatchConfirmResponse `json:"alibaba_wdk_ums_retrieve_batch_confirm_response,omitempty"` 
    AlibabaWdkUmsRetrieveBatchConfirmResponse
}

/* model for simplify = false
type AlibabaWdkUmsRetrieveBatchConfirmResponse struct {

    // result
    
    Result  *struct {
        UtmsResult  *UtmsResult `json:"utms_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkUmsRetrieveBatchConfirmResponse struct {

    // result
    Result   *UtmsResult `json:"result,omitempty"`

}
