package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
充值退款 APIResponse
alibaba.alsc.crm.recharge.uncharge.update

充值退款
*/
type AlibabaAlscCrmRechargeUnchargeUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeUnchargeUpdateResponse
}

type AlibabaAlscCrmRechargeUnchargeUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_uncharge_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
