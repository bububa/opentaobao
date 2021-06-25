package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖品添加接口 APIRequest
alibaba.marketing.lottery.award.append

抽奖平台奖品添加接口，目前仅用于奖池众筹项目
*/
type AlibabaMarketingLotteryAwardAppendRequest struct {
    model.Params

    // 奖品添加请求对象
    lotteryAwardAppend   *LotteryAwardAppendDto 

}

func NewAlibabaMarketingLotteryAwardAppendRequest() *AlibabaMarketingLotteryAwardAppendRequest{
    return &AlibabaMarketingLotteryAwardAppendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMarketingLotteryAwardAppendRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.award.append"
}

func (r AlibabaMarketingLotteryAwardAppendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMarketingLotteryAwardAppendRequest) SetLotteryAwardAppend(lotteryAwardAppend *LotteryAwardAppendDto) error {
    r.lotteryAwardAppend = lotteryAwardAppend
    r.Set("lottery_award_append", lotteryAwardAppend)
    return nil
}

func (r AlibabaMarketingLotteryAwardAppendRequest) GetLotteryAwardAppend() *LotteryAwardAppendDto {
    return r.lotteryAwardAppend
}

