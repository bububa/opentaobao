package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherWatchAPIRequest 监控预约数据 API请求
// taobao.game.deliveryvoucher.watch
//
// 监控预约数据
type TaobaoGameDeliveryvoucherWatchAPIRequest struct {
	model.Params
	// 入参
	_param0 *WatchAppointmentRequest
}

// NewTaobaoGameDeliveryvoucherWatchRequest 初始化TaobaoGameDeliveryvoucherWatchAPIRequest对象
func NewTaobaoGameDeliveryvoucherWatchRequest() *TaobaoGameDeliveryvoucherWatchAPIRequest {
	return &TaobaoGameDeliveryvoucherWatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherWatchAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.watch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherWatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoGameDeliveryvoucherWatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *TaobaoGameDeliveryvoucherWatchAPIRequest) SetParam0(_param0 *WatchAppointmentRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoGameDeliveryvoucherWatchAPIRequest) GetParam0() *WatchAppointmentRequest {
	return r._param0
}
