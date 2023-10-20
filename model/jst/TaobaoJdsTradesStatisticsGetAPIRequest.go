package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojdstradesstatisticsgetAPIRequest 获取订单数量统计结果 API请求
// taobao.jds.trades.statistics.get
//
// 获取订单数量统计结果
type TaobaojdstradesstatisticsgetAPIRequest struct {
	model.Params
	// 查询的日期，格式如YYYYMMDD的日期对应的数字
	_date int64
}

// NewTaobaojdstradesstatisticsgetRequest 初始化TaobaojdstradesstatisticsgetAPIRequest对象
func NewTaobaojdstradesstatisticsgetRequest() *TaobaojdstradesstatisticsgetAPIRequest {
	return &TaobaojdstradesstatisticsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojdstradesstatisticsgetAPIRequest) GetApiMethodName() string {
	return "taobao.jds.trades.statistics.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojdstradesstatisticsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojdstradesstatisticsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDate is Date Setter
// 查询的日期，格式如YYYYMMDD的日期对应的数字
func (r *TaobaojdstradesstatisticsgetAPIRequest) SetDate(_date int64) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r TaobaojdstradesstatisticsgetAPIRequest) GetDate() int64 {
	return r._date
}
