package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id) APIResponse
alibaba.alihealth.drug.kyt.dr.getentinfo

根据企业名称查询ID
*/
type AlibabaAlihealthDrugKytDrGetentinfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDrGetentinfoResponse
}

type AlibabaAlihealthDrugKytDrGetentinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_getentinfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytDrGetentinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
