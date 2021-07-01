package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillIiUpdateAPIRequest
电子面单云打印更新接口 API请求
cainiao.waybill.ii.update

商家更新电子面单号对应的面单信息。 */
type CainiaoWaybillIiUpdateAPIRequest struct {
	model.Params
	// 更新请求信息
	_paramWaybillCloudPrintUpdateRequest *WaybillCloudPrintUpdateRequest
}

// NewCainiaoWaybillIiUpdateRequest 初始化CainiaoWaybillIiUpdateAPIRequest对象
func NewCainiaoWaybillIiUpdateRequest() *CainiaoWaybillIiUpdateAPIRequest {
	return &CainiaoWaybillIiUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiUpdateAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamWaybillCloudPrintUpdateRequest Setter
// 更新请求信息
func (r *CainiaoWaybillIiUpdateAPIRequest) SetParamWaybillCloudPrintUpdateRequest(_paramWaybillCloudPrintUpdateRequest *WaybillCloudPrintUpdateRequest) error {
	r._paramWaybillCloudPrintUpdateRequest = _paramWaybillCloudPrintUpdateRequest
	r.Set("param_waybill_cloud_print_update_request", _paramWaybillCloudPrintUpdateRequest)
	return nil
}

// Get ParamWaybillCloudPrintUpdateRequest Getter
func (r CainiaoWaybillIiUpdateAPIRequest) GetParamWaybillCloudPrintUpdateRequest() *WaybillCloudPrintUpdateRequest {
	return r._paramWaybillCloudPrintUpdateRequest
}
