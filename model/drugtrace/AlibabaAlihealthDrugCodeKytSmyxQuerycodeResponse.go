package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
扫码营销码查询 APIResponse
alibaba.alihealth.drug.code.kyt.smyx.querycode

此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
核查平台优先过滤非8开头的，长度非20位数字的码信息。
*/
type AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugCodeKytSmyxQuerycodeResponse
}

type AlibabaAlihealthDrugCodeKytSmyxQuerycodeResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_smyx_querycode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 最外层结果
    
    Result   *AlibabaAlihealthDrugCodeKytSmyxQuerycodeResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
