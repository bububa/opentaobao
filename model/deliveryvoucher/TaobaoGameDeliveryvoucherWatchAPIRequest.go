package deliveryvoucher

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoGameDeliveryvoucherWatchAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
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

var poolTaobaoGameDeliveryvoucherWatchAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoGameDeliveryvoucherWatchRequest()
	},
}

// GetTaobaoGameDeliveryvoucherWatchRequest 从 sync.Pool 获取 TaobaoGameDeliveryvoucherWatchAPIRequest
func GetTaobaoGameDeliveryvoucherWatchAPIRequest() *TaobaoGameDeliveryvoucherWatchAPIRequest {
	return poolTaobaoGameDeliveryvoucherWatchAPIRequest.Get().(*TaobaoGameDeliveryvoucherWatchAPIRequest)
}

// ReleaseTaobaoGameDeliveryvoucherWatchAPIRequest 将 TaobaoGameDeliveryvoucherWatchAPIRequest 放入 sync.Pool
func ReleaseTaobaoGameDeliveryvoucherWatchAPIRequest(v *TaobaoGameDeliveryvoucherWatchAPIRequest) {
	v.Reset()
	poolTaobaoGameDeliveryvoucherWatchAPIRequest.Put(v)
}
