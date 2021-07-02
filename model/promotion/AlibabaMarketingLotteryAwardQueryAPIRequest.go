package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMarketingLotteryAwardQueryAPIRequest 抽奖平台查询可用奖品接口 API请求
// alibaba.marketing.lottery.award.query
//
// 抽奖平台查询可用奖品接口
type AlibabaMarketingLotteryAwardQueryAPIRequest struct {
	model.Params
	// 查询奖品请求对象
	_lotteryAwardInstQuery *LotteryAwardInstQueryDto
}

// NewAlibabaMarketingLotteryAwardQueryRequest 初始化AlibabaMarketingLotteryAwardQueryAPIRequest对象
func NewAlibabaMarketingLotteryAwardQueryRequest() *AlibabaMarketingLotteryAwardQueryAPIRequest {
	return &AlibabaMarketingLotteryAwardQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryAwardQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.marketing.lottery.award.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryAwardQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LotteryAwardInstQuery Setter
// 查询奖品请求对象
func (r *AlibabaMarketingLotteryAwardQueryAPIRequest) SetLotteryAwardInstQuery(_lotteryAwardInstQuery *LotteryAwardInstQueryDto) error {
	r._lotteryAwardInstQuery = _lotteryAwardInstQuery
	r.Set("lottery_award_inst_query", _lotteryAwardInstQuery)
	return nil
}

// Get LotteryAwardInstQuery Getter
func (r AlibabaMarketingLotteryAwardQueryAPIRequest) GetLotteryAwardInstQuery() *LotteryAwardInstQueryDto {
	return r._lotteryAwardInstQuery
}
