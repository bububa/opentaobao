package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池绑定接口 API请求
alibaba.marketing.lottery.activity.bind

抽奖平台奖池关联接口
*/
type AlibabaMarketingLotteryActivityBindRequest struct {
    model.Params
    // 关联抽奖活动请求对象
    _lotteryActivityRel   *LotteryActivityRelDto
}

// 初始化AlibabaMarketingLotteryActivityBindRequest对象
func NewAlibabaMarketingLotteryActivityBindRequest() *AlibabaMarketingLotteryActivityBindRequest{
    return &AlibabaMarketingLotteryActivityBindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryActivityBindRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.activity.bind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryActivityBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LotteryActivityRel Setter
// 关联抽奖活动请求对象
func (r *AlibabaMarketingLotteryActivityBindRequest) SetLotteryActivityRel(_lotteryActivityRel *LotteryActivityRelDto) error {
    r._lotteryActivityRel = _lotteryActivityRel
    r.Set("lottery_activity_rel", _lotteryActivityRel)
    return nil
}

// LotteryActivityRel Getter
func (r AlibabaMarketingLotteryActivityBindRequest) GetLotteryActivityRel() *LotteryActivityRelDto {
    return r._lotteryActivityRel
}
