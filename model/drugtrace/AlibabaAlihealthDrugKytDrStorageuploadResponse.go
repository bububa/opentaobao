package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗存储温度上传 APIResponse
alibaba.alihealth.drug.kyt.dr.storageupload

疫苗存储温度上传
*/
type AlibabaAlihealthDrugKytDrStorageuploadAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDrStorageuploadResponse
}

type AlibabaAlihealthDrugKytDrStorageuploadResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_storageupload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugKytDrStorageuploadResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
