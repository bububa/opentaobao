package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillIPrintAPIRequest 打印确认接口v1.0 API请求
// taobao.wlb.waybill.i.print
//
// 打印面单前的校验接口，判断面单号信息与订单信息是否匹配。
type TaobaoWlbWaybillIPrintAPIRequest struct {
	model.Params
	// 打印请求
	_waybillApplyPrintCheckRequest *WaybillApplyPrintCheckRequest
}

// NewTaobaoWlbWaybillIPrintRequest 初始化TaobaoWlbWaybillIPrintAPIRequest对象
func NewTaobaoWlbWaybillIPrintRequest() *TaobaoWlbWaybillIPrintAPIRequest {
	return &TaobaoWlbWaybillIPrintAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbWaybillIPrintAPIRequest) Reset() {
	r._waybillApplyPrintCheckRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillIPrintAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.print"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillIPrintAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbWaybillIPrintAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillApplyPrintCheckRequest is WaybillApplyPrintCheckRequest Setter
// 打印请求
func (r *TaobaoWlbWaybillIPrintAPIRequest) SetWaybillApplyPrintCheckRequest(_waybillApplyPrintCheckRequest *WaybillApplyPrintCheckRequest) error {
	r._waybillApplyPrintCheckRequest = _waybillApplyPrintCheckRequest
	r.Set("waybill_apply_print_check_request", _waybillApplyPrintCheckRequest)
	return nil
}

// GetWaybillApplyPrintCheckRequest WaybillApplyPrintCheckRequest Getter
func (r TaobaoWlbWaybillIPrintAPIRequest) GetWaybillApplyPrintCheckRequest() *WaybillApplyPrintCheckRequest {
	return r._waybillApplyPrintCheckRequest
}

var poolTaobaoWlbWaybillIPrintAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbWaybillIPrintRequest()
	},
}

// GetTaobaoWlbWaybillIPrintRequest 从 sync.Pool 获取 TaobaoWlbWaybillIPrintAPIRequest
func GetTaobaoWlbWaybillIPrintAPIRequest() *TaobaoWlbWaybillIPrintAPIRequest {
	return poolTaobaoWlbWaybillIPrintAPIRequest.Get().(*TaobaoWlbWaybillIPrintAPIRequest)
}

// ReleaseTaobaoWlbWaybillIPrintAPIRequest 将 TaobaoWlbWaybillIPrintAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbWaybillIPrintAPIRequest(v *TaobaoWlbWaybillIPrintAPIRequest) {
	v.Reset()
	poolTaobaoWlbWaybillIPrintAPIRequest.Put(v)
}
