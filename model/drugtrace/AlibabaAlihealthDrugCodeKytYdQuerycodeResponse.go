package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询追溯码对应的药品信息（药店） APIResponse
alibaba.alihealth.drug.code.kyt.yd.querycode

此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
核查平台优先过滤非8开头的，长度非20位数字的码信息。
*/
type AlibabaAlihealthDrugCodeKytYdQuerycodeAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugCodeKytYdQuerycodeResponse
}

type AlibabaAlihealthDrugCodeKytYdQuerycodeResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_yd_querycode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 最外层结果
    
    Result   *AlibabaAlihealthDrugCodeKytYdQuerycodeResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
