package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
中药片关联关系上传 API返回值 
alibaba.alihealth.zy.uploadrelation

中药片关联关系上传
*/
type AlibabaAlihealthZyUploadrelationAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthZyUploadrelationResponse
}

// 中药片关联关系上传 成功返回结果
type AlibabaAlihealthZyUploadrelationResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_zy_uploadrelation_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaAlihealthZyUploadrelationResultModel `json:"result,omitempty" xml:"result,omitempty"`
}