package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
储值消费退款(逆向) APIResponse
alibaba.alsc.crm.recharge.undedut.update

新增储值消费退款接口
*/
type AlibabaAlscCrmRechargeUndedutUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeUndedutUpdateResponse
}

type AlibabaAlscCrmRechargeUndedutUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_undedut_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
