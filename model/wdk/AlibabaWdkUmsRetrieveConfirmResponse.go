package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
回流单－外部对已拉取到的UMS单据进行确认 APIResponse
alibaba.wdk.ums.retrieve.confirm

回流单－外部对已拉取到的UMS单据进行确认
*/
type AlibabaWdkUmsRetrieveConfirmAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkUmsRetrieveConfirmResponse `json:"alibaba_wdk_ums_retrieve_confirm_response,omitempty"`
}

type AlibabaWdkUmsRetrieveConfirmResponse struct {

    // result
    Result   *UtmsResult `json:"result,omitempty"`

}
