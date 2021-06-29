package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增往来单位企业 APIResponse
alibaba.alihealth.drug.kyt.saveent

新增往来单位企业记录
*/
type AlibabaAlihealthDrugKytSaveentAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytSaveentResponse
}

type AlibabaAlihealthDrugKytSaveentResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_saveent_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 往来单位新增接口返回
    
    Result   *AlibabaAlihealthDrugKytSaveentResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
