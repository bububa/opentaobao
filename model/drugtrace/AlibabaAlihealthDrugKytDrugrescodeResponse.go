package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询药品码段信息 APIResponse
alibaba.alihealth.drug.kyt.drugrescode

查询药品码段信息
*/
type AlibabaAlihealthDrugKytDrugrescodeAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDrugrescodeResponse
}

type AlibabaAlihealthDrugKytDrugrescodeResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_drugrescode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytDrugrescodeResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
