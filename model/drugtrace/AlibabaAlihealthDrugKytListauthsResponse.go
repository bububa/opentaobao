package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
企业搜索自己授权的物流企业 APIResponse
alibaba.alihealth.drug.kyt.listauths

企业搜索自己授权的物流企业
*/
type AlibabaAlihealthDrugKytListauthsAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytListauthsResponse
}

type AlibabaAlihealthDrugKytListauthsResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_listauths_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytListauthsResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
