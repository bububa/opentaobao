package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
零售药店往来单位新增 APIResponse
alibaba.alihealth.drug.lsyd.saveent

新增往来单位企业记录
*/
type AlibabaAlihealthDrugLsydSaveentAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugLsydSaveentResponse
}

type AlibabaAlihealthDrugLsydSaveentResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_lsyd_saveent_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 往来单位新增接口返回
    
    Result   *AlibabaAlihealthDrugLsydSaveentResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
