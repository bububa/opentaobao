package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
码查询药品 APIResponse
alibaba.alihealth.drug.kyt.querydruginfo

通过追溯码查询药品信息
*/
type AlibabaAlihealthDrugKytQuerydruginfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytQuerydruginfoResponse
}

type AlibabaAlihealthDrugKytQuerydruginfoResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_querydruginfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaAlihealthDrugKytQuerydruginfoResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
