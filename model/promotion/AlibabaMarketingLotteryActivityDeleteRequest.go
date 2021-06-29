package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台活动删除接口 API请求
alibaba.marketing.lottery.activity.delete

抽奖平台活动删除接口
*/
type AlibabaMarketingLotteryActivityDeleteRequest struct {
    model.Params
    // 抽奖活动删除对象
    lotteryActivityDelete   *LotteryActivityDeleteDto
}

// 初始化AlibabaMarketingLotteryActivityDeleteRequest对象
func NewAlibabaMarketingLotteryActivityDeleteRequest() *AlibabaMarketingLotteryActivityDeleteRequest{
    return &AlibabaMarketingLotteryActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryActivityDeleteRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.activity.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LotteryActivityDelete Setter
// 抽奖活动删除对象
func (r *AlibabaMarketingLotteryActivityDeleteRequest) SetLotteryActivityDelete(lotteryActivityDelete *LotteryActivityDeleteDto) error {
    r.lotteryActivityDelete = lotteryActivityDelete
    r.Set("lottery_activity_delete", lotteryActivityDelete)
    return nil
}

// LotteryActivityDelete Getter
func (r AlibabaMarketingLotteryActivityDeleteRequest) GetLotteryActivityDelete() *LotteryActivityDeleteDto {
    return r.lotteryActivityDelete
}
