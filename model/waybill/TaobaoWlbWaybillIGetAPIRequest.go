package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillIGetAPIRequest 获取物流服务商电子面单号v1.0 API请求
// taobao.wlb.waybill.i.get
//
// 商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。
type TaobaoWlbWaybillIGetAPIRequest struct {
	model.Params
	// 面单申请
	_waybillApplyNewRequest *WaybillApplyNewRequest
}

// NewTaobaoWlbWaybillIGetRequest 初始化TaobaoWlbWaybillIGetAPIRequest对象
func NewTaobaoWlbWaybillIGetRequest() *TaobaoWlbWaybillIGetAPIRequest {
	return &TaobaoWlbWaybillIGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbWaybillIGetAPIRequest) Reset() {
	r._waybillApplyNewRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillIGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillIGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbWaybillIGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillApplyNewRequest is WaybillApplyNewRequest Setter
// 面单申请
func (r *TaobaoWlbWaybillIGetAPIRequest) SetWaybillApplyNewRequest(_waybillApplyNewRequest *WaybillApplyNewRequest) error {
	r._waybillApplyNewRequest = _waybillApplyNewRequest
	r.Set("waybill_apply_new_request", _waybillApplyNewRequest)
	return nil
}

// GetWaybillApplyNewRequest WaybillApplyNewRequest Getter
func (r TaobaoWlbWaybillIGetAPIRequest) GetWaybillApplyNewRequest() *WaybillApplyNewRequest {
	return r._waybillApplyNewRequest
}

var poolTaobaoWlbWaybillIGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbWaybillIGetRequest()
	},
}

// GetTaobaoWlbWaybillIGetRequest 从 sync.Pool 获取 TaobaoWlbWaybillIGetAPIRequest
func GetTaobaoWlbWaybillIGetAPIRequest() *TaobaoWlbWaybillIGetAPIRequest {
	return poolTaobaoWlbWaybillIGetAPIRequest.Get().(*TaobaoWlbWaybillIGetAPIRequest)
}

// ReleaseTaobaoWlbWaybillIGetAPIRequest 将 TaobaoWlbWaybillIGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbWaybillIGetAPIRequest(v *TaobaoWlbWaybillIGetAPIRequest) {
	v.Reset()
	poolTaobaoWlbWaybillIGetAPIRequest.Put(v)
}
