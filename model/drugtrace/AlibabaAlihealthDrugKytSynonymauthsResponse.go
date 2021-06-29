package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流企业查询货主企业信息 APIResponse
alibaba.alihealth.drug.kyt.synonymauths

物流企业查询货主企业信息
*/
type AlibabaAlihealthDrugKytSynonymauthsAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytSynonymauthsResponse
}

type AlibabaAlihealthDrugKytSynonymauthsResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_synonymauths_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytSynonymauthsResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
