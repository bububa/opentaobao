package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池创建接口 APIResponse
alibaba.marketing.lottery.activity.create

抽奖平台奖池创建接口
*/
type AlibabaMarketingLotteryActivityCreateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMarketingLotteryActivityCreateResponse `json:"alibaba_marketing_lottery_activity_create_response,omitempty"`
}

type AlibabaMarketingLotteryActivityCreateResponse struct {

    // 错误码
    MsgCode   int64 `json:"msg_code,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty"`

    // 抽奖活动
    LotteryActivity   *LotteryActivityExtendDto `json:"lottery_activity,omitempty"`

}
