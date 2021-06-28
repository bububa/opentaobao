package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台抽奖规则保存接口 APIResponse
alibaba.marketing.lottery.rule.save

抽奖平台抽奖规则保存接口，对于同一主体，保存新规则会失效老的规则
*/
type AlibabaMarketingLotteryRuleSaveAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_marketing_lottery_rule_save_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 保存成功
    
    Result   bool `json:"result,omitempty" xml:"