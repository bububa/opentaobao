package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询药品详细信息 APIResponse
alibaba.alihealth.drug.kyt.drugdetail

查询药品详细信息
*/
type AlibabaAlihealthDrugKytDrugdetailAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDrugdetailResponse
}

type AlibabaAlihealthDrugKytDrugdetailResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_drugdetail_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytDrugdetailResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
