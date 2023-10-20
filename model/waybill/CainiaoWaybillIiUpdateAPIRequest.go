package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiUpdateAPIRequest 电子面单云打印更新接口 API请求
// cainiao.waybill.ii.update
//
// 商家更新电子面单号对应的面单信息。
type CainiaoWaybillIiUpdateAPIRequest struct {
	model.Params
	// 更新请求信息
	_paramWaybillCloudPrintUpdateRequest *WaybillCloudPrintUpdateRequest
}

// NewCainiaoWaybillIiUpdateRequest 初始化CainiaoWaybillIiUpdateAPIRequest对象
func NewCainiaoWaybillIiUpdateRequest() *CainiaoWaybillIiUpdateAPIRequest {
	return &CainiaoWaybillIiUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoWaybillIiUpdateAPIRequest) Reset() {
	r._paramWaybillCloudPrintUpdateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiUpdateAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoWaybillIiUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamWaybillCloudPrintUpdateRequest is ParamWaybillCloudPrintUpdateRequest Setter
// 更新请求信息
func (r *CainiaoWaybillIiUpdateAPIRequest) SetParamWaybillCloudPrintUpdateRequest(_paramWaybillCloudPrintUpdateRequest *WaybillCloudPrintUpdateRequest) error {
	r._paramWaybillCloudPrintUpdateRequest = _paramWaybillCloudPrintUpdateRequest
	r.Set("param_waybill_cloud_print_update_request", _paramWaybillCloudPrintUpdateRequest)
	return nil
}

// GetParamWaybillCloudPrintUpdateRequest ParamWaybillCloudPrintUpdateRequest Getter
func (r CainiaoWaybillIiUpdateAPIRequest) GetParamWaybillCloudPrintUpdateRequest() *WaybillCloudPrintUpdateRequest {
	return r._paramWaybillCloudPrintUpdateRequest
}

var poolCainiaoWaybillIiUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoWaybillIiUpdateRequest()
	},
}

// GetCainiaoWaybillIiUpdateRequest 从 sync.Pool 获取 CainiaoWaybillIiUpdateAPIRequest
func GetCainiaoWaybillIiUpdateAPIRequest() *CainiaoWaybillIiUpdateAPIRequest {
	return poolCainiaoWaybillIiUpdateAPIRequest.Get().(*CainiaoWaybillIiUpdateAPIRequest)
}

// ReleaseCainiaoWaybillIiUpdateAPIRequest 将 CainiaoWaybillIiUpdateAPIRequest 放入 sync.Pool
func ReleaseCainiaoWaybillIiUpdateAPIRequest(v *CainiaoWaybillIiUpdateAPIRequest) {
	v.Reset()
	poolCainiaoWaybillIiUpdateAPIRequest.Put(v)
}
