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
    AlibabaWdkUmsRetrieveBatchConfirmResponse
}

type AlibabaWdkUmsRetrieveBatchConfirmResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_ums_retrieve_batch_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
