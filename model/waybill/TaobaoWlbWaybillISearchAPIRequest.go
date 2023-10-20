package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwaybillisearchAPIRequest 查询面单服务订购及面单使用情况v1.0 API请求
// taobao.wlb.waybill.i.search
//
// 获取发货地&amp;CP开通状态&amp;账户的使用情况
type TaobaowlbwaybillisearchAPIRequest struct {
	model.Params
	// 查询网点信息
	_waybillApplyRequest *WaybillApplyRequest
}

// NewTaobaowlbwaybillisearchRequest 初始化TaobaowlbwaybillisearchAPIRequest对象
func NewTaobaowlbwaybillisearchRequest() *TaobaowlbwaybillisearchAPIRequest {
	return &TaobaowlbwaybillisearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwaybillisearchAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwaybillisearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwaybillisearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillApplyRequest is WaybillApplyRequest Setter
// 查询网点信息
func (r *TaobaowlbwaybillisearchAPIRequest) SetWaybillApplyRequest(_waybillApplyRequest *WaybillApplyRequest) error {
	r._waybillApplyRequest = _waybillApplyRequest
	r.Set("waybill_apply_request", _waybillApplyRequest)
	return nil
}

// GetWaybillApplyRequest WaybillApplyRequest Getter
func (r TaobaowlbwaybillisearchAPIRequest) GetWaybillApplyRequest() *WaybillApplyRequest {
	return r._waybillApplyRequest
}
