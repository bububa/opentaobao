package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
多融通过企业ID得到一个企业的详细信息 APIResponse
alibaba.alihealth.drug.kyt.dr.getbyentid

根据企业主键查看企业详细信息
*/
type AlibabaAlihealthDrugKytDrGetbyentidAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDrGetbyentidResponse
}

type AlibabaAlihealthDrugKytDrGetbyentidResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_getbyentid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytDrGetbyentidResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
