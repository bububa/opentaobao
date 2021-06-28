package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
储值消费 APIResponse
alibaba.alsc.crm.recharge.dedut.update

增加储值消费接口
*/
type AlibabaAlscCrmRechargeDedutUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeDedutUpdateResponse
}

type AlibabaAlscCrmRechargeDedutUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_dedut_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
