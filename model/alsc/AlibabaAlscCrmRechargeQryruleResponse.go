package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
储值规则下行 APIResponse
alibaba.alsc.crm.recharge.qryrule

储值规则下行
*/
type AlibabaAlscCrmRechargeQryruleAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeQryruleResponse
}

type AlibabaAlscCrmRechargeQryruleResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_qryrule_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回模型
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
