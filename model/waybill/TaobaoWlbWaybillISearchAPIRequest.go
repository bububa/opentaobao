package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillISearchAPIRequest 查询面单服务订购及面单使用情况v1.0 API请求
// taobao.wlb.waybill.i.search
//
// 获取发货地&amp;CP开通状态&amp;账户的使用情况
type TaobaoWlbWaybillISearchAPIRequest struct {
	model.Params
	// 查询网点信息
	_waybillApplyRequest *WaybillApplyRequest
}

// NewTaobaoWlbWaybillISearchRequest 初始化TaobaoWlbWaybillISearchAPIRequest对象
func NewTaobaoWlbWaybillISearchRequest() *TaobaoWlbWaybillISearchAPIRequest {
	return &TaobaoWlbWaybillISearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillISearchAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillISearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetWaybillApplyRequest is WaybillApplyRequest Setter
// 查询网点信息
func (r *TaobaoWlbWaybillISearchAPIRequest) SetWaybillApplyRequest(_waybillApplyRequest *WaybillApplyRequest) error {
	r._waybillApplyRequest = _waybillApplyRequest
	r.Set("waybill_apply_request", _waybillApplyRequest)
	return nil
}

// GetWaybillApplyRequest WaybillApplyRequest Getter
func (r TaobaoWlbWaybillISearchAPIRequest) GetWaybillApplyRequest() *WaybillApplyRequest {
	return r._waybillApplyRequest
}
