package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗运输温度上传 APIResponse
alibaba.alihealth.drug.kyt.dr.transportupload

疫苗运输温度上传
*/
type AlibabaAlihealthDrugKytDrTransportuploadAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDrTransportuploadResponse
}

type AlibabaAlihealthDrugKytDrTransportuploadResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_transportupload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugKytDrTransportuploadResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
