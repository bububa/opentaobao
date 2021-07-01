package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMarketingLotteryActivityBindAPIRequest
抽奖平台奖池绑定接口 API请求
alibaba.marketing.lottery.activity.bind

抽奖平台奖池关联接口 */
type AlibabaMarketingLotteryActivityBindAPIRequest struct {
	model.Params
	// 关联抽奖活动请求对象
	_lotteryActivityRel *LotteryActivityRelDto
}

// NewAlibabaMarketingLotteryActivityBindRequest 初始化AlibabaMarketingLotteryActivityBindAPIRequest对象
func NewAlibabaMarketingLotteryActivityBindRequest() *AlibabaMarketingLotteryActivityBindAPIRequest {
	return &AlibabaMarketingLotteryActivityBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryActivityBindAPIRequest) GetApiMethodName() string {
	return "alibaba.marketing.lottery.activity.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryActivityBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LotteryActivityRel Setter
// 关联抽奖活动请求对象
func (r *AlibabaMarketingLotteryActivityBindAPIRequest) SetLotteryActivityRel(_lotteryActivityRel *LotteryActivityRelDto) error {
	r._lotteryActivityRel = _lotteryActivityRel
	r.Set("lottery_activity_rel", _lotteryActivityRel)
	return nil
}

// Get LotteryActivityRel Getter
func (r AlibabaMarketingLotteryActivityBindAPIRequest) GetLotteryActivityRel() *LotteryActivityRelDto {
	return r._lotteryActivityRel
}
