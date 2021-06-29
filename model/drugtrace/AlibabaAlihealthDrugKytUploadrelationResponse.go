package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关联关系上传 APIResponse
alibaba.alihealth.drug.kyt.uploadrelation

关联关系上传
*/
type AlibabaAlihealthDrugKytUploadrelationAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytUploadrelationResponse
}

type AlibabaAlihealthDrugKytUploadrelationResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_uploadrelation_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugKytUploadrelationResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
