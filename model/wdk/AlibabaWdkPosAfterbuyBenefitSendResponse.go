package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
生态pos购后发放权益 APIResponse
alibaba.wdk.pos.afterbuy.benefit.send

生态pos购后发放权益接口开放
*/
type AlibabaWdkPosAfterbuyBenefitSendAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkPosAfterbuyBenefitSendResponse `json:"alibaba_wdk_pos_afterbuy_benefit_send_response,omitempty"`
}

type AlibabaWdkPosAfterbuyBenefitSendResponse struct {

    // 返回结果
    Result   *BmResult `json:"result,omitempty"`

}
