package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJdsTradesStatisticsGetAPIRequest 获取订单数量统计结果 API请求
// taobao.jds.trades.statistics.get
//
// 获取订单数量统计结果
type TaobaoJdsTradesStatisticsGetAPIRequest struct {
	model.Params
	// 查询的日期，格式如YYYYMMDD的日期对应的数字
	_date int64
}

// NewTaobaoJdsTradesStatisticsGetRequest 初始化TaobaoJdsTradesStatisticsGetAPIRequest对象
func NewTaobaoJdsTradesStatisticsGetRequest() *TaobaoJdsTradesStatisticsGetAPIRequest {
	return &TaobaoJdsTradesStatisticsGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJdsTradesStatisticsGetAPIRequest) Reset() {
	r._date = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJdsTradesStatisticsGetAPIRequest) GetApiMethodName() string {
	return "taobao.jds.trades.statistics.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJdsTradesStatisticsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJdsTradesStatisticsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDate is Date Setter
// 查询的日期，格式如YYYYMMDD的日期对应的数字
func (r *TaobaoJdsTradesStatisticsGetAPIRequest) SetDate(_date int64) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r TaobaoJdsTradesStatisticsGetAPIRequest) GetDate() int64 {
	return r._date
}

var poolTaobaoJdsTradesStatisticsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJdsTradesStatisticsGetRequest()
	},
}

// GetTaobaoJdsTradesStatisticsGetRequest 从 sync.Pool 获取 TaobaoJdsTradesStatisticsGetAPIRequest
func GetTaobaoJdsTradesStatisticsGetAPIRequest() *TaobaoJdsTradesStatisticsGetAPIRequest {
	return poolTaobaoJdsTradesStatisticsGetAPIRequest.Get().(*TaobaoJdsTradesStatisticsGetAPIRequest)
}

// ReleaseTaobaoJdsTradesStatisticsGetAPIRequest 将 TaobaoJdsTradesStatisticsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoJdsTradesStatisticsGetAPIRequest(v *TaobaoJdsTradesStatisticsGetAPIRequest) {
	v.Reset()
	poolTaobaoJdsTradesStatisticsGetAPIRequest.Put(v)
}
