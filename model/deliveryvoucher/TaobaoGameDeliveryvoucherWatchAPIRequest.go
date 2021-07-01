package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoGameDeliveryvoucherWatchAPIRequest
监控预约数据 API请求
taobao.game.deliveryvoucher.watch

监控预约数据 */
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
func (r TaobaoGameDeliveryvoucherWatchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 入参
func (r *TaobaoGameDeliveryvoucherWatchAPIRequest) SetParam0(_param0 *WatchAppointmentRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoGameDeliveryvoucherWatchAPIRequest) GetParam0() *WatchAppointmentRequest {
	return r._param0
}
