package film

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFilmLotteryPerformanceAPIRequest) Reset() {
	r._lotteryPerformanceTopParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFilmLotteryPerformanceAPIRequest) GetApiMethodName() string {
	return "taobao.film.lottery.performance"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFilmLotteryPerformanceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFilmLotteryPerformanceAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoFilmLotteryPerformanceAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFilmLotteryPerformanceRequest()
	},
}

// GetTaobaoFilmLotteryPerformanceRequest 从 sync.Pool 获取 TaobaoFilmLotteryPerformanceAPIRequest
func GetTaobaoFilmLotteryPerformanceAPIRequest() *TaobaoFilmLotteryPerformanceAPIRequest {
	return poolTaobaoFilmLotteryPerformanceAPIRequest.Get().(*TaobaoFilmLotteryPerformanceAPIRequest)
}

// ReleaseTaobaoFilmLotteryPerformanceAPIRequest 将 TaobaoFilmLotteryPerformanceAPIRequest 放入 sync.Pool
func ReleaseTaobaoFilmLotteryPerformanceAPIRequest(v *TaobaoFilmLotteryPerformanceAPIRequest) {
	v.Reset()
	poolTaobaoFilmLotteryPerformanceAPIRequest.Put(v)
}
