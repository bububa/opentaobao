package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
中药片关联关系上传 APIResponse
alibaba.alihealth.zy.uploadrelation

中药片关联关系上传
*/
type AlibabaAlihealthZyUploadrelationAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthZyUploadrelationResponse
}

type AlibabaAlihealthZyUploadrelationResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_zy_uploadrelation_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthZyUploadrelationResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
