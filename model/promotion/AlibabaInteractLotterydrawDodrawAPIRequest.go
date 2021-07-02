package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractLotterydrawDodrawAPIRequest 无线端抽奖接口 API请求
// alibaba.interact.lotterydraw.dodraw
//
// 商家抽奖平台无线端抽奖接口开放
type AlibabaInteractLotterydrawDodrawAPIRequest struct {
	model.Params
	// 抽奖请求对象
	_lotteryDrawQuery *LotteryDrawQueryDto
}

// NewAlibabaInteractLotterydrawDodrawRequest 初始化AlibabaInteractLotterydrawDodrawAPIRequest对象
func NewAlibabaInteractLotterydrawDodrawRequest() *AlibabaInteractLotterydrawDodrawAPIRequest {
	return &AlibabaInteractLotterydrawDodrawAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractLotterydrawDodrawAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.lotterydraw.dodraw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractLotterydrawDodrawAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LotteryDrawQuery Setter
// 抽奖请求对象
func (r *AlibabaInteractLotterydrawDodrawAPIRequest) SetLotteryDrawQuery(_lotteryDrawQuery *LotteryDrawQueryDto) error {
	r._lotteryDrawQuery = _lotteryDrawQuery
	r.Set("lottery_draw_query", _lotteryDrawQuery)
	return nil
}

// Get LotteryDrawQuery Getter
func (r AlibabaInteractLotterydrawDodrawAPIRequest) GetLotteryDrawQuery() *LotteryDrawQueryDto {
	return r._lotteryDrawQuery
}
