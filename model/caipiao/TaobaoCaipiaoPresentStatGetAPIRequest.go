package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocaipiaopresentstatgetAPIRequest 获取卖家按天统计的彩票赠送数据 API请求
// taobao.caipiao.present.stat.get
//
// 查询卖家一段时间内按天统计的彩票赠送数据，只支持查询90天以内的数据.
type TaobaocaipiaopresentstatgetAPIRequest struct {
	model.Params
	// 指定查询的天数，从当前日期（不包括当前日期）向前推算的天数，可为空。如果为空、0、负数或者大于90天，则设置为默认的90天。举例：当天是20120703， days=2， 则统计数据的日期为：20120702，20120701.
	_days int64
}

// NewTaobaocaipiaopresentstatgetRequest 初始化TaobaocaipiaopresentstatgetAPIRequest对象
func NewTaobaocaipiaopresentstatgetRequest() *TaobaocaipiaopresentstatgetAPIRequest {
	return &TaobaocaipiaopresentstatgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocaipiaopresentstatgetAPIRequest) GetApiMethodName() string {
	return "taobao.caipiao.present.stat.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocaipiaopresentstatgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocaipiaopresentstatgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDays is Days Setter
// 指定查询的天数，从当前日期（不包括当前日期）向前推算的天数，可为空。如果为空、0、负数或者大于90天，则设置为默认的90天。举例：当天是20120703， days=2， 则统计数据的日期为：20120702，20120701.
func (r *TaobaocaipiaopresentstatgetAPIRequest) SetDays(_days int64) error {
	r._days = _days
	r.Set("days", _days)
	return nil
}

// GetDays Days Getter
func (r TaobaocaipiaopresentstatgetAPIRequest) GetDays() int64 {
	return r._days
}
