package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
标准激活卡 APIResponse
alibaba.alsc.crm.card.active

激活卡
*/
type AlibabaAlscCrmCardActiveAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCardActiveResponse
}

type AlibabaAlscCrmCardActiveResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_card_active_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
