package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMarketingLotteryActivityUnbindAPIRequest 抽奖平台奖池解绑接口 API请求
// alibaba.marketing.lottery.activity.unbind
//
// 抽奖平台奖池解绑接口
type AlibabaMarketingLotteryActivityUnbindAPIRequest struct {
	model.Params
	// 解绑抽奖活动请求对象
	_lotteryActivityRel *LotteryActivityRelDto
}

// NewAlibabaMarketingLotteryActivityUnbindRequest 初始化AlibabaMarketingLotteryActivityUnbindAPIRequest对象
func NewAlibabaMarketingLotteryActivityUnbindRequest() *AlibabaMarketingLotteryActivityUnbindAPIRequest {
	return &AlibabaMarketingLotteryActivityUnbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryActivityUnbindAPIRequest) GetApiMethodName() string {
	return "alibaba.marketing.lottery.activity.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryActivityUnbindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LotteryActivityRel Setter
// 解绑抽奖活动请求对象
func (r *AlibabaMarketingLotteryActivityUnbindAPIRequest) SetLotteryActivityRel(_lotteryActivityRel *LotteryActivityRelDto) error {
	r._lotteryActivityRel = _lotteryActivityRel
	r.Set("lottery_activity_rel", _lotteryActivityRel)
	return nil
}

// Get LotteryActivityRel Getter
func (r AlibabaMarketingLotteryActivityUnbindAPIRequest) GetLotteryActivityRel() *LotteryActivityRelDto {
	return r._lotteryActivityRel
}
