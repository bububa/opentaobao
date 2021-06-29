package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流代货主查往来单位接口 APIResponse
alibaba.alihealth.drug.kyt.listparts.byagent

代理企业查询往来单位列表
*/
type AlibabaAlihealthDrugKytListpartsByagentAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytListpartsByagentResponse
}

type AlibabaAlihealthDrugKytListpartsByagentResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_listparts_byagent_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytListpartsByagentResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
