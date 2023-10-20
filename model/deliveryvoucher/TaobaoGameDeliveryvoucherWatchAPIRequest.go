package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaogamedeliveryvoucherwatchAPIRequest 监控预约数据 API请求
// taobao.game.deliveryvoucher.watch
//
// 监控预约数据
type TaobaogamedeliveryvoucherwatchAPIRequest struct {
	model.Params
	// 入参
	_param0 *WatchAppointmentRequest
}

// NewTaobaogamedeliveryvoucherwatchRequest 初始化TaobaogamedeliveryvoucherwatchAPIRequest对象
func NewTaobaogamedeliveryvoucherwatchRequest() *TaobaogamedeliveryvoucherwatchAPIRequest {
	return &TaobaogamedeliveryvoucherwatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaogamedeliveryvoucherwatchAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.watch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaogamedeliveryvoucherwatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaogamedeliveryvoucherwatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *TaobaogamedeliveryvoucherwatchAPIRequest) SetParam0(_param0 *WatchAppointmentRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaogamedeliveryvoucherwatchAPIRequest) GetParam0() *WatchAppointmentRequest {
	return r._param0
}
