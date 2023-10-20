package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwaybillicancelAPIRequest 商家取消获取的电子面单号v1.0 API请求
// taobao.wlb.waybill.i.cancel
//
// 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type TaobaowlbwaybillicancelAPIRequest struct {
	model.Params
	// 取消接口入参
	_waybillApplyCancelRequest *WaybillApplyCancelRequest
}

// NewTaobaowlbwaybillicancelRequest 初始化TaobaowlbwaybillicancelAPIRequest对象
func NewTaobaowlbwaybillicancelRequest() *TaobaowlbwaybillicancelAPIRequest {
	return &TaobaowlbwaybillicancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwaybillicancelAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwaybillicancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwaybillicancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillApplyCancelRequest is WaybillApplyCancelRequest Setter
// 取消接口入参
func (r *TaobaowlbwaybillicancelAPIRequest) SetWaybillApplyCancelRequest(_waybillApplyCancelRequest *WaybillApplyCancelRequest) error {
	r._waybillApplyCancelRequest = _waybillApplyCancelRequest
	r.Set("waybill_apply_cancel_request", _waybillApplyCancelRequest)
	return nil
}

// GetWaybillApplyCancelRequest WaybillApplyCancelRequest Getter
func (r TaobaowlbwaybillicancelAPIRequest) GetWaybillApplyCancelRequest() *WaybillApplyCancelRequest {
	return r._waybillApplyCancelRequest
}
