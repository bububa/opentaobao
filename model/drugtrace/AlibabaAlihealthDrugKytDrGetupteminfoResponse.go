package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取上游温度信息（疫苗） APIResponse
alibaba.alihealth.drug.kyt.dr.getupteminfo

根据追溯码及企业ID获取上游运输及存储温度信息（疫苗）
*/
type AlibabaAlihealthDrugKytDrGetupteminfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDrGetupteminfoResponse
}

type AlibabaAlihealthDrugKytDrGetupteminfoResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_getupteminfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugKytDrGetupteminfoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
