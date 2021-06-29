package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 APIResponse
alibaba.alihealth.drug.kyt.getentinfo

根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
*/
type AlibabaAlihealthDrugKytGetentinfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytGetentinfoResponse
}

type AlibabaAlihealthDrugKytGetentinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_getentinfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytGetentinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
