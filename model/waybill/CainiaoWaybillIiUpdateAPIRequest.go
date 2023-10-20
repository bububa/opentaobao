package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybilliiupdateAPIRequest 电子面单云打印更新接口 API请求
// cainiao.waybill.ii.update
//
// 商家更新电子面单号对应的面单信息。
type CainiaowaybilliiupdateAPIRequest struct {
	model.Params
	// 更新请求信息
	_paramWaybillCloudPrintUpdateRequest *WaybillCloudPrintUpdateRequest
}

// NewCainiaowaybilliiupdateRequest 初始化CainiaowaybilliiupdateAPIRequest对象
func NewCainiaowaybilliiupdateRequest() *CainiaowaybilliiupdateAPIRequest {
	return &CainiaowaybilliiupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaowaybilliiupdateAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaowaybilliiupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaowaybilliiupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamWaybillCloudPrintUpdateRequest is ParamWaybillCloudPrintUpdateRequest Setter
// 更新请求信息
func (r *CainiaowaybilliiupdateAPIRequest) SetParamWaybillCloudPrintUpdateRequest(_paramWaybillCloudPrintUpdateRequest *WaybillCloudPrintUpdateRequest) error {
	r._paramWaybillCloudPrintUpdateRequest = _paramWaybillCloudPrintUpdateRequest
	r.Set("param_waybill_cloud_print_update_request", _paramWaybillCloudPrintUpdateRequest)
	return nil
}

// GetParamWaybillCloudPrintUpdateRequest ParamWaybillCloudPrintUpdateRequest Getter
func (r CainiaowaybilliiupdateAPIRequest) GetParamWaybillCloudPrintUpdateRequest() *WaybillCloudPrintUpdateRequest {
	return r._paramWaybillCloudPrintUpdateRequest
}
