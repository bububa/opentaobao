package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
绑定物理卡 APIResponse
alibaba.alsc.crm.card.bindcard

将会员卡和实例物理卡绑定在一起
*/
type AlibabaAlscCrmCardBindcardAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alsc_crm_card_bindcard_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"