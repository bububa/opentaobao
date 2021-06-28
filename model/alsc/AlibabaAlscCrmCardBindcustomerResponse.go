package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卡号绑定顾客 APIResponse
alibaba.alsc.crm.card.bindcustomer

为卡号绑定顾客
*/
type AlibabaAlscCrmCardBindcustomerAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCardBindcustomerResponse
}

type AlibabaAlscCrmCardBindcustomerResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_card_bindcustomer_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
