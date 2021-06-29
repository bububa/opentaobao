package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
储值充值 APIResponse
alibaba.alsc.crm.recharge.charge.update

顾客储值账户充值
*/
type AlibabaAlscCrmRechargeChargeUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeChargeUpdateResponse
}

type AlibabaAlscCrmRechargeChargeUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_charge_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
