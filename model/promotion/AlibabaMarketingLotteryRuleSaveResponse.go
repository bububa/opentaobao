package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台抽奖规则保存接口 APIResponse
alibaba.marketing.lottery.rule.save

抽奖平台抽奖规则保存接口，对于同一主体，保存新规则会失效老的规则
*/
type AlibabaMarketingLotteryRuleSaveAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMarketingLotteryRuleSaveResponse `json:"alibaba_marketing_lottery_rule_save_response,omitempty"`
}

type AlibabaMarketingLotteryRuleSaveResponse struct {

    // 保存成功
    Result   bool `json:"result,omitempty"`

    // code
    MsgCode   int64 `json:"msg_code,omitempty"`

    // 接口调用成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // msg
    MsgInfo   string `json:"msg_info,omitempty"`

}
