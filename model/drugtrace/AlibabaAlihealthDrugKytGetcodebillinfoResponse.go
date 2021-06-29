package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据码获取基本和单据信息 APIResponse
alibaba.alihealth.drug.kyt.getcodebillinfo

根据码信息获取基本信息和单据信息
*/
type AlibabaAlihealthDrugKytGetcodebillinfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytGetcodebillinfoResponse
}

type AlibabaAlihealthDrugKytGetcodebillinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_getcodebillinfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 和三方交互最外层model对象
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
