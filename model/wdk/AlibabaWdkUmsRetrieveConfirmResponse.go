package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
回流单－外部对已拉取到的UMS单据进行确认 APIResponse
alibaba.wdk.ums.retrieve.confirm

回流单－外部对已拉取到的UMS单据进行确认
*/
type AlibabaWdkUmsRetrieveConfirmAPIResponse struct {
    model.CommonResponse
    AlibabaWdkUmsRetrieveConfirmResponse
}

type AlibabaWdkUmsRetrieveConfirmResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_ums_retrieve_confirm_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
