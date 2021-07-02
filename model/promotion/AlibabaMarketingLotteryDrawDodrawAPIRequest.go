package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMarketingLotteryDrawDodrawAPIRequest 抽奖平台抽奖接口 API请求
// alibaba.marketing.lottery.draw.dodraw
//
// 抽奖平台PC端抽奖接口
type AlibabaMarketingLotteryDrawDodrawAPIRequest struct {
	model.Params
	// 抽奖请求对象
	_lotteryDrawQuery *LotteryDrawQueryDto
}

// NewAlibabaMarketingLotteryDrawDodrawRequest 初始化AlibabaMarketingLotteryDrawDodrawAPIRequest对象
func NewAlibabaMarketingLotteryDrawDodrawRequest() *AlibabaMarketingLotteryDrawDodrawAPIRequest {
	return &AlibabaMarketingLotteryDrawDodrawAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryDrawDodrawAPIRequest) GetApiMethodName() string {
	return "alibaba.marketing.lottery.draw.dodraw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryDrawDodrawAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LotteryDrawQuery Setter
// 抽奖请求对象
func (r *AlibabaMarketingLotteryDrawDodrawAPIRequest) SetLotteryDrawQuery(_lotteryDrawQuery *LotteryDrawQueryDto) error {
	r._lotteryDrawQuery = _lotteryDrawQuery
	r.Set("lottery_draw_query", _lotteryDrawQuery)
	return nil
}

// Get LotteryDrawQuery Getter
func (r AlibabaMarketingLotteryDrawDodrawAPIRequest) GetLotteryDrawQuery() *LotteryDrawQueryDto {
	return r._lotteryDrawQuery
}
