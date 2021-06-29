package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询药品目录信息 APIResponse
alibaba.alihealth.drug.kyt.drugtable

查询药品目录信息
*/
type AlibabaAlihealthDrugKytDrugtableAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDrugtableResponse
}

type AlibabaAlihealthDrugKytDrugtableResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_drugtable_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytDrugtableResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
