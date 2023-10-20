package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwaybilligetAPIRequest 获取物流服务商电子面单号v1.0 API请求
// taobao.wlb.waybill.i.get
//
// 商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。
type TaobaowlbwaybilligetAPIRequest struct {
	model.Params
	// 面单申请
	_waybillApplyNewRequest *WaybillApplyNewRequest
}

// NewTaobaowlbwaybilligetRequest 初始化TaobaowlbwaybilligetAPIRequest对象
func NewTaobaowlbwaybilligetRequest() *TaobaowlbwaybilligetAPIRequest {
	return &TaobaowlbwaybilligetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwaybilligetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwaybilligetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwaybilligetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillApplyNewRequest is WaybillApplyNewRequest Setter
// 面单申请
func (r *TaobaowlbwaybilligetAPIRequest) SetWaybillApplyNewRequest(_waybillApplyNewRequest *WaybillApplyNewRequest) error {
	r._waybillApplyNewRequest = _waybillApplyNewRequest
	r.Set("waybill_apply_new_request", _waybillApplyNewRequest)
	return nil
}

// GetWaybillApplyNewRequest WaybillApplyNewRequest Getter
func (r TaobaowlbwaybilligetAPIRequest) GetWaybillApplyNewRequest() *WaybillApplyNewRequest {
	return r._waybillApplyNewRequest
}
