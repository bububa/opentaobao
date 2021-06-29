package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户测评结果回传接口 APIResponse
alibaba.alihealth.alipaypfm.assess.result.sync

用户测评结果回传接口
*/
type AlibabaAlihealthAlipaypfmAssessResultSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthAlipaypfmAssessResultSyncResponse
}

type AlibabaAlihealthAlipaypfmAssessResultSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_alipaypfm_assess_result_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 和三方交互最外层model对象
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
