package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMarketingLotteryActivityCreateAPIRequest 抽奖平台奖池创建接口 API请求
// alibaba.marketing.lottery.activity.create
//
// 抽奖平台奖池创建接口
type AlibabaMarketingLotteryActivityCreateAPIRequest struct {
	model.Params
	// 抽奖活动创建请求对象
	_lotteryActivityCreate *LotteryActivityCreateDto
}

// NewAlibabaMarketingLotteryActivityCreateRequest 初始化AlibabaMarketingLotteryActivityCreateAPIRequest对象
func NewAlibabaMarketingLotteryActivityCreateRequest() *AlibabaMarketingLotteryActivityCreateAPIRequest {
	return &AlibabaMarketingLotteryActivityCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryActivityCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.marketing.lottery.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryActivityCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LotteryActivityCreate Setter
// 抽奖活动创建请求对象
func (r *AlibabaMarketingLotteryActivityCreateAPIRequest) SetLotteryActivityCreate(_lotteryActivityCreate *LotteryActivityCreateDto) error {
	r._lotteryActivityCreate = _lotteryActivityCreate
	r.Set("lottery_activity_create", _lotteryActivityCreate)
	return nil
}

// Get LotteryActivityCreate Getter
func (r AlibabaMarketingLotteryActivityCreateAPIRequest) GetLotteryActivityCreate() *LotteryActivityCreateDto {
	return r._lotteryActivityCreate
}
