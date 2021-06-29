package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
标准开卡流程 APIResponse
alibaba.alsc.crm.card.open

标准开卡流程
*/
type AlibabaAlscCrmCardOpenAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCardOpenResponse
}

type AlibabaAlscCrmCardOpenResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_card_open_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
