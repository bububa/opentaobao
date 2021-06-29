package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
储值账户退充值校验 APIResponse
alibaba.alsc.crm.recharge.unchargecheck.get

储值账户退充值校验接口
*/
type AlibabaAlscCrmRechargeUnchargecheckGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeUnchargecheckGetResponse
}

type AlibabaAlscCrmRechargeUnchargecheckGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_unchargecheck_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
