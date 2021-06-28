package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
互动到店抽奖 APIResponse
alibaba.interact.isvlottery.idraw

互动到店抽奖
*/
type AlibabaInteractIsvlotteryIdrawAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractIsvlotteryIdrawResponse `json:"alibaba_interact_isvlottery_idraw_response,omitempty"` 
    AlibabaInteractIsvlotteryIdrawResponse
}

/* model for simplify = false
type AlibabaInteractIsvlotteryIdrawResponse struct {

    // 抽奖中奖信息
    
    Data  *struct {
        LotteryProxyResult  *LotteryProxyResult `json:"lottery_proxy_result,omitempty"`
    } `json:"data,omitempty"`
    

    // 是否调用成功
    
    Succ   bool `json:"succ,omitempty"`
    

    // 错误信息
    
    BizCode   string `json:"biz_code,omitempty"`
    

    // 错误信息描述
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

type AlibabaInteractIsvlotteryIdrawResponse struct {

    // 抽奖中奖信息
    Data   *LotteryProxyResult `json:"data,omitempty"`

    // 是否调用成功
    Succ   bool `json:"succ,omitempty"`

    // 错误信息
    BizCode   string `json:"biz_code,omitempty"`

    // 错误信息描述
    ErrorMsg   string `json:"error_msg,omitempty"`

}
