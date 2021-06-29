package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
储值规则下行 API返回值 
alibaba.alsc.crm.recharge.qryrule

储值规则下行
*/
type AlibabaAlscCrmRechargeQryruleAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeQryruleResponse
}

// 储值规则下行 成功返回结果
type AlibabaAlscCrmRechargeQryruleResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_qryrule_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回模型
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
