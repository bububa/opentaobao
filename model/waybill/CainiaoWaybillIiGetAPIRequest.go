package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiGetAPIRequest 电子面单云打印接口 API请求
// cainiao.waybill.ii.get
//
// 菜鸟电子面单的云打印申请电子面单号的方法
type CainiaoWaybillIiGetAPIRequest struct {
	model.Params
	// 入参信息
	_paramWaybillCloudPrintApplyNewRequest *WaybillCloudPrintApplyNewRequest
}

// NewCainiaoWaybillIiGetRequest 初始化CainiaoWaybillIiGetAPIRequest对象
func NewCainiaoWaybillIiGetRequest() *CainiaoWaybillIiGetAPIRequest {
	return &CainiaoWaybillIiGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoWaybillIiGetAPIRequest) Reset() {
	r._paramWaybillCloudPrintApplyNewRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiGetAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoWaybillIiGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamWaybillCloudPrintApplyNewRequest is ParamWaybillCloudPrintApplyNewRequest Setter
// 入参信息
func (r *CainiaoWaybillIiGetAPIRequest) SetParamWaybillCloudPrintApplyNewRequest(_paramWaybillCloudPrintApplyNewRequest *WaybillCloudPrintApplyNewRequest) error {
	r._paramWaybillCloudPrintApplyNewRequest = _paramWaybillCloudPrintApplyNewRequest
	r.Set("param_waybill_cloud_print_apply_new_request", _paramWaybillCloudPrintApplyNewRequest)
	return nil
}

// GetParamWaybillCloudPrintApplyNewRequest ParamWaybillCloudPrintApplyNewRequest Getter
func (r CainiaoWaybillIiGetAPIRequest) GetParamWaybillCloudPrintApplyNewRequest() *WaybillCloudPrintApplyNewRequest {
	return r._paramWaybillCloudPrintApplyNewRequest
}

var poolCainiaoWaybillIiGetAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoWaybillIiGetRequest()
	},
}

// GetCainiaoWaybillIiGetRequest 从 sync.Pool 获取 CainiaoWaybillIiGetAPIRequest
func GetCainiaoWaybillIiGetAPIRequest() *CainiaoWaybillIiGetAPIRequest {
	return poolCainiaoWaybillIiGetAPIRequest.Get().(*CainiaoWaybillIiGetAPIRequest)
}

// ReleaseCainiaoWaybillIiGetAPIRequest 将 CainiaoWaybillIiGetAPIRequest 放入 sync.Pool
func ReleaseCainiaoWaybillIiGetAPIRequest(v *CainiaoWaybillIiGetAPIRequest) {
	v.Reset()
	poolCainiaoWaybillIiGetAPIRequest.Put(v)
}
