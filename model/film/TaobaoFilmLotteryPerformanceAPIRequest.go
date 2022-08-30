package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmLotteryPerformanceAPIRequest 淘票票履约发放权益 API请求
// taobao.film.lottery.performance
//
// 对外第三方合作渠道通过抽奖形式发放权益
type TaobaoFilmLotteryPerformanceAPIRequest struct {
	model.Params
	// 入参
	_lotteryPerformanceTopParam *LotteryPerformanceTopParam
}

// NewTaobaoFilmLotteryPerformanceRequest 初始化TaobaoFilmLotteryPerformanceAPIRequest对象
func NewTaobaoFilmLotteryPerformanceRequest() *TaobaoFilmLotteryPerformanceAPIRequest {
	return &TaobaoFilmLotteryPerformanceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFilmLotteryPerformanceAPIRequest) GetApiMethodName() string {
	return "taobao.film.lottery.performance"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFilmLotteryPerformanceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLotteryPerformanceTopParam is LotteryPerformanceTopParam Setter
// 入参
func (r *TaobaoFilmLotteryPerformanceAPIRequest) SetLotteryPerformanceTopParam(_lotteryPerformanceTopParam *LotteryPerformanceTopParam) error {
	r._lotteryPerformanceTopParam = _lotteryPerformanceTopParam
	r.Set("lottery_performance_top_param", _lotteryPerformanceTopParam)
	return nil
}

// GetLotteryPerformanceTopParam LotteryPerformanceTopParam Getter
func (r TaobaoFilmLotteryPerformanceAPIRequest) GetLotteryPerformanceTopParam() *LotteryPerformanceTopParam {
	return r._lotteryPerformanceTopParam
}
