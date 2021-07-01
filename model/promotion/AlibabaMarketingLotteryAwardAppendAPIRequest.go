package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMarketingLotteryAwardAppendAPIRequest
抽奖平台奖品添加接口 API请求
alibaba.marketing.lottery.award.append

抽奖平台奖品添加接口，目前仅用于奖池众筹项目 */
type AlibabaMarketingLotteryAwardAppendAPIRequest struct {
	model.Params
	// 奖品添加请求对象
	_lotteryAwardAppend *LotteryAwardAppendDto
}

// NewAlibabaMarketingLotteryAwardAppendRequest 初始化AlibabaMarketingLotteryAwardAppendAPIRequest对象
func NewAlibabaMarketingLotteryAwardAppendRequest() *AlibabaMarketingLotteryAwardAppendAPIRequest {
	return &AlibabaMarketingLotteryAwardAppendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryAwardAppendAPIRequest) GetApiMethodName() string {
	return "alibaba.marketing.lottery.award.append"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryAwardAppendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LotteryAwardAppend Setter
// 奖品添加请求对象
func (r *AlibabaMarketingLotteryAwardAppendAPIRequest) SetLotteryAwardAppend(_lotteryAwardAppend *LotteryAwardAppendDto) error {
	r._lotteryAwardAppend = _lotteryAwardAppend
	r.Set("lottery_award_append", _lotteryAwardAppend)
	return nil
}

// Get LotteryAwardAppend Getter
func (r AlibabaMarketingLotteryAwardAppendAPIRequest) GetLotteryAwardAppend() *LotteryAwardAppendDto {
	return r._lotteryAwardAppend
}
