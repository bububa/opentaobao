package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwaybillifullupdateAPIRequest 面单信息更新接口v1.0 API请求
// taobao.wlb.waybill.i.fullupdate
//
// 商家更新电子面单号对应的订单信息。&lt;br/&gt;&lt;br/&gt;a.涉及到订单信息（如拆/合单、取消订单等）、修改发货地的逻辑时候，需要使用CANCEL+GET组合操作。一般的面单信息更新使用该接口；&lt;br/&gt;&lt;br/&gt;b.该接口是全量修改接口，对于不做更新的字段也要把原有的字段值传进去，否则做为修改为空处理。
type TaobaowlbwaybillifullupdateAPIRequest struct {
	model.Params
	// 更新面单信息请求
	_waybillApplyFullUpdateRequest *WaybillApplyFullUpdateRequest
}

// NewTaobaowlbwaybillifullupdateRequest 初始化TaobaowlbwaybillifullupdateAPIRequest对象
func NewTaobaowlbwaybillifullupdateRequest() *TaobaowlbwaybillifullupdateAPIRequest {
	return &TaobaowlbwaybillifullupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwaybillifullupdateAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.fullupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwaybillifullupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwaybillifullupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillApplyFullUpdateRequest is WaybillApplyFullUpdateRequest Setter
// 更新面单信息请求
func (r *TaobaowlbwaybillifullupdateAPIRequest) SetWaybillApplyFullUpdateRequest(_waybillApplyFullUpdateRequest *WaybillApplyFullUpdateRequest) error {
	r._waybillApplyFullUpdateRequest = _waybillApplyFullUpdateRequest
	r.Set("waybill_apply_full_update_request", _waybillApplyFullUpdateRequest)
	return nil
}

// GetWaybillApplyFullUpdateRequest WaybillApplyFullUpdateRequest Getter
func (r TaobaowlbwaybillifullupdateAPIRequest) GetWaybillApplyFullUpdateRequest() *WaybillApplyFullUpdateRequest {
	return r._waybillApplyFullUpdateRequest
}
