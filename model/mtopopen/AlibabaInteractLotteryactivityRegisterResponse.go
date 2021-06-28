package mtopopen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
回传抽奖相关参数 APIResponse
alibaba.interact.lotteryactivity.register

提供接口供三方应用将数据回传到平台
*/
type AlibabaInteractLotteryactivityRegisterAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractLotteryactivityRegisterResponse `json:"alibaba_interact_lotteryactivity_register_response,omitempty"` 
    AlibabaInteractLotteryactivityRegisterResponse
}

/* model for simplify = false
type AlibabaInteractLotteryactivityRegisterResponse struct {

    // 接口返回model
    
    Result  *struct {
        AlibabaInteractLotteryactivityRegisterResult  *AlibabaInteractLotteryactivityRegisterResult `json:"alibaba_interact_lotteryactivity_register_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaInteractLotteryactivityRegisterResponse struct {

    // 接口返回model
    Result   *AlibabaInteractLotteryactivityRegisterResult `json:"result,omitempty"`

}
