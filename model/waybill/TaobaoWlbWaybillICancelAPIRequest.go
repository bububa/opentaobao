package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillICancelAPIRequest 商家取消获取的电子面单号v1.0 API请求
// taobao.wlb.waybill.i.cancel
//
// 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type TaobaoWlbWaybillICancelAPIRequest struct {
	model.Params
	// 取消接口入参
	_waybillApplyCancelRequest *WaybillApplyCancelRequest
}

// NewTaobaoWlbWaybillICancelRequest 初始化TaobaoWlbWaybillICancelAPIRequest对象
func NewTaobaoWlbWaybillICancelRequest() *TaobaoWlbWaybillICancelAPIRequest {
	return &TaobaoWlbWaybillICancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillICancelAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillICancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WaybillApplyCancelRequest Setter
// 取消接口入参
func (r *TaobaoWlbWaybillICancelAPIRequest) SetWaybillApplyCancelRequest(_waybillApplyCancelRequest *WaybillApplyCancelRequest) error {
	r._waybillApplyCancelRequest = _waybillApplyCancelRequest
	r.Set("waybill_apply_cancel_request", _waybillApplyCancelRequest)
	return nil
}

// Get WaybillApplyCancelRequest Getter
func (r TaobaoWlbWaybillICancelAPIRequest) GetWaybillApplyCancelRequest() *WaybillApplyCancelRequest {
	return r._waybillApplyCancelRequest
}
