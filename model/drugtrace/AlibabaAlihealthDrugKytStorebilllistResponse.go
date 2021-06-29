package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
零售端平台单据查询 APIResponse
alibaba.alihealth.drug.kyt.storebilllist

零售端平台单据查询
*/
type AlibabaAlihealthDrugKytStorebilllistAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytStorebilllistResponse
}

type AlibabaAlihealthDrugKytStorebilllistResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_storebilllist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytStorebilllistResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
