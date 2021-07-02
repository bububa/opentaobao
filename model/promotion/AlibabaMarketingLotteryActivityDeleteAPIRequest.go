package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMarketingLotteryActivityDeleteAPIRequest 抽奖平台活动删除接口 API请求
// alibaba.marketing.lottery.activity.delete
//
// 抽奖平台活动删除接口
type AlibabaMarketingLotteryActivityDeleteAPIRequest struct {
	model.Params
	// 抽奖活动删除对象
	_lotteryActivityDelete *LotteryActivityDeleteDto
}

// NewAlibabaMarketingLotteryActivityDeleteRequest 初始化AlibabaMarketingLotteryActivityDeleteAPIRequest对象
func NewAlibabaMarketingLotteryActivityDeleteRequest() *AlibabaMarketingLotteryActivityDeleteAPIRequest {
	return &AlibabaMarketingLotteryActivityDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryActivityDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.marketing.lottery.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryActivityDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LotteryActivityDelete Setter
// 抽奖活动删除对象
func (r *AlibabaMarketingLotteryActivityDeleteAPIRequest) SetLotteryActivityDelete(_lotteryActivityDelete *LotteryActivityDeleteDto) error {
	r._lotteryActivityDelete = _lotteryActivityDelete
	r.Set("lottery_activity_delete", _lotteryActivityDelete)
	return nil
}

// Get LotteryActivityDelete Getter
func (r AlibabaMarketingLotteryActivityDeleteAPIRequest) GetLotteryActivityDelete() *LotteryActivityDeleteDto {
	return r._lotteryActivityDelete
}
