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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alsc_crm_recharge_dedut_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"