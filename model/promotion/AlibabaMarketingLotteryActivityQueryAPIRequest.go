package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMarketingLotteryActivityQueryAPIRequest 抽奖平台奖池查询接口 API请求
// alibaba.marketing.lottery.activity.query
//
// 抽奖平台奖池查询接口
type AlibabaMarketingLotteryActivityQueryAPIRequest struct {
	model.Params
	// 查询抽奖活动请求对象
	_lotteryActivityQuery *LotteryActivityQueryDto
}

// NewAlibabaMarketingLotteryActivityQueryRequest 初始化AlibabaMarketingLotteryActivityQueryAPIRequest对象
func NewAlibabaMarketingLotteryActivityQueryRequest() *AlibabaMarketingLotteryActivityQueryAPIRequest {
	return &AlibabaMarketingLotteryActivityQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryActivityQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.marketing.lottery.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryActivityQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LotteryActivityQuery Setter
// 查询抽奖活动请求对象
func (r *AlibabaMarketingLotteryActivityQueryAPIRequest) SetLotteryActivityQuery(_lotteryActivityQuery *LotteryActivityQueryDto) error {
	r._lotteryActivityQuery = _lotteryActivityQuery
	r.Set("lottery_activity_query", _lotteryActivityQuery)
	return nil
}

// Get LotteryActivityQuery Getter
func (r AlibabaMarketingLotteryActivityQueryAPIRequest) GetLotteryActivityQuery() *LotteryActivityQueryDto {
	return r._lotteryActivityQuery
}
