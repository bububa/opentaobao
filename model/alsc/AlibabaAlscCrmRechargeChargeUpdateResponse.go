package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
储值充值 API返回值 
alibaba.alsc.crm.recharge.charge.update

顾客储值账户充值
*/
type AlibabaAlscCrmRechargeChargeUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeChargeUpdateResponse
}

// 储值充值 成功返回结果
type AlibabaAlscCrmRechargeChargeUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_charge_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口结果
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
