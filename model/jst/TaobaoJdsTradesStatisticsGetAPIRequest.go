package jst

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJdsTradesStatisticsGetAPIRequest) GetApiMethodName() string {
	return "taobao.jds.trades.statistics.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJdsTradesStatisticsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Date Setter
// 查询的日期，格式如YYYYMMDD的日期对应的数字
func (r *TaobaoJdsTradesStatisticsGetAPIRequest) SetDate(_date int64) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// Get Date Getter
func (r TaobaoJdsTradesStatisticsGetAPIRequest) GetDate() int64 {
	return r._date
}
