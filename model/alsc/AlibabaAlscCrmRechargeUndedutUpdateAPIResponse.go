package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
储值消费退款(逆向) API返回值 
alibaba.alsc.crm.recharge.undedut.update

新增储值消费退款接口
*/
type AlibabaAlscCrmRechargeUndedutUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeUndedutUpdateAPIResponseModel
}

// 储值消费退款(逆向) 成功返回结果
type AlibabaAlscCrmRechargeUndedutUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_undedut_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口结果
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
