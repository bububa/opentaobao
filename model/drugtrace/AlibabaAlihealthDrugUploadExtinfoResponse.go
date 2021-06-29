package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
中药饮片及器械对接 APIResponse
alibaba.alihealth.drug.upload.extinfo

中药饮片及器械对接
*/
type AlibabaAlihealthDrugUploadExtinfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugUploadExtinfoResponse
}

type AlibabaAlihealthDrugUploadExtinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_upload_extinfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaAlihealthDrugUploadExtinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
