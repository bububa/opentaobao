package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwaybilliprintAPIRequest 打印确认接口v1.0 API请求
// taobao.wlb.waybill.i.print
//
// 打印面单前的校验接口，判断面单号信息与订单信息是否匹配。
type TaobaowlbwaybilliprintAPIRequest struct {
	model.Params
	// 打印请求
	_waybillApplyPrintCheckRequest *WaybillApplyPrintCheckRequest
}

// NewTaobaowlbwaybilliprintRequest 初始化TaobaowlbwaybilliprintAPIRequest对象
func NewTaobaowlbwaybilliprintRequest() *TaobaowlbwaybilliprintAPIRequest {
	return &TaobaowlbwaybilliprintAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwaybilliprintAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.print"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwaybilliprintAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwaybilliprintAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillApplyPrintCheckRequest is WaybillApplyPrintCheckRequest Setter
// 打印请求
func (r *TaobaowlbwaybilliprintAPIRequest) SetWaybillApplyPrintCheckRequest(_waybillApplyPrintCheckRequest *WaybillApplyPrintCheckRequest) error {
	r._waybillApplyPrintCheckRequest = _waybillApplyPrintCheckRequest
	r.Set("waybill_apply_print_check_request", _waybillApplyPrintCheckRequest)
	return nil
}

// GetWaybillApplyPrintCheckRequest WaybillApplyPrintCheckRequest Getter
func (r TaobaowlbwaybilliprintAPIRequest) GetWaybillApplyPrintCheckRequest() *WaybillApplyPrintCheckRequest {
	return r._waybillApplyPrintCheckRequest
}
