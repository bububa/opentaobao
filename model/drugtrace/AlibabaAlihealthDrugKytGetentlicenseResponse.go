package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取企业资质 APIResponse
alibaba.alihealth.drug.kyt.getentlicense

获取企业的资质信息。
*/
type AlibabaAlihealthDrugKytGetentlicenseAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytGetentlicenseResponse
}

type AlibabaAlihealthDrugKytGetentlicenseResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_getentlicense_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytGetentlicenseResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
