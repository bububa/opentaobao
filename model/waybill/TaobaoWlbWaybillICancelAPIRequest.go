package waybill

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbWaybillICancelAPIRequest) Reset() {
	r._waybillApplyCancelRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillICancelAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillICancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbWaybillICancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillApplyCancelRequest is WaybillApplyCancelRequest Setter
// 取消接口入参
func (r *TaobaoWlbWaybillICancelAPIRequest) SetWaybillApplyCancelRequest(_waybillApplyCancelRequest *WaybillApplyCancelRequest) error {
	r._waybillApplyCancelRequest = _waybillApplyCancelRequest
	r.Set("waybill_apply_cancel_request", _waybillApplyCancelRequest)
	return nil
}

// GetWaybillApplyCancelRequest WaybillApplyCancelRequest Getter
func (r TaobaoWlbWaybillICancelAPIRequest) GetWaybillApplyCancelRequest() *WaybillApplyCancelRequest {
	return r._waybillApplyCancelRequest
}

var poolTaobaoWlbWaybillICancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbWaybillICancelRequest()
	},
}

// GetTaobaoWlbWaybillICancelRequest 从 sync.Pool 获取 TaobaoWlbWaybillICancelAPIRequest
func GetTaobaoWlbWaybillICancelAPIRequest() *TaobaoWlbWaybillICancelAPIRequest {
	return poolTaobaoWlbWaybillICancelAPIRequest.Get().(*TaobaoWlbWaybillICancelAPIRequest)
}

// ReleaseTaobaoWlbWaybillICancelAPIRequest 将 TaobaoWlbWaybillICancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbWaybillICancelAPIRequest(v *TaobaoWlbWaybillICancelAPIRequest) {
	v.Reset()
	poolTaobaoWlbWaybillICancelAPIRequest.Put(v)
}
