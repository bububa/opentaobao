package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
储值账户充值前校验 API返回值 
alibaba.alsc.crm.recharge.chargeprecheck.get

储值账户充值前校验接口
*/
type AlibabaAlscCrmRechargeChargeprecheckGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeChargeprecheckGetAPIResponseModel
}

// 储值账户充值前校验 成功返回结果
type AlibabaAlscCrmRechargeChargeprecheckGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_chargeprecheck_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口结果
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
