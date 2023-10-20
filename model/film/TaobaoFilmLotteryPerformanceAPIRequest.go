package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofilmlotteryperformanceAPIRequest 淘票票履约发放权益 API请求
// taobao.film.lottery.performance
//
// 对外第三方合作渠道通过抽奖形式发放权益
type TaobaofilmlotteryperformanceAPIRequest struct {
	model.Params
	// 入参
	_lotteryPerformanceTopParam *LotteryPerformanceTopParam
}

// NewTaobaofilmlotteryperformanceRequest 初始化TaobaofilmlotteryperformanceAPIRequest对象
func NewTaobaofilmlotteryperformanceRequest() *TaobaofilmlotteryperformanceAPIRequest {
	return &TaobaofilmlotteryperformanceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofilmlotteryperformanceAPIRequest) GetApiMethodName() string {
	return "taobao.film.lottery.performance"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofilmlotteryperformanceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofilmlotteryperformanceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLotteryPerformanceTopParam is LotteryPerformanceTopParam Setter
// 入参
func (r *TaobaofilmlotteryperformanceAPIRequest) SetLotteryPerformanceTopParam(_lotteryPerformanceTopParam *LotteryPerformanceTopParam) error {
	r._lotteryPerformanceTopParam = _lotteryPerformanceTopParam
	r.Set("lottery_performance_top_param", _lotteryPerformanceTopParam)
	return nil
}

// GetLotteryPerformanceTopParam LotteryPerformanceTopParam Getter
func (r TaobaofilmlotteryperformanceAPIRequest) GetLotteryPerformanceTopParam() *LotteryPerformanceTopParam {
	return r._lotteryPerformanceTopParam
}
