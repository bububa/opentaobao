package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询积分流水 APIResponse
alibaba.alsc.crm.point.querypointflow

分页查询积分流水
*/
type AlibabaAlscCrmPointQuerypointflowAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmPointQuerypointflowResponse
}

type AlibabaAlscCrmPointQuerypointflowResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_point_querypointflow_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 分页返回模型
    
    Result   *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
