package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybilliigetAPIRequest 电子面单云打印接口 API请求
// cainiao.waybill.ii.get
//
// 菜鸟电子面单的云打印申请电子面单号的方法
type CainiaowaybilliigetAPIRequest struct {
	model.Params
	// 入参信息
	_paramWaybillCloudPrintApplyNewRequest *WaybillCloudPrintApplyNewRequest
}

// NewCainiaowaybilliigetRequest 初始化CainiaowaybilliigetAPIRequest对象
func NewCainiaowaybilliigetRequest() *CainiaowaybilliigetAPIRequest {
	return &CainiaowaybilliigetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaowaybilliigetAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaowaybilliigetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaowaybilliigetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamWaybillCloudPrintApplyNewRequest is ParamWaybillCloudPrintApplyNewRequest Setter
// 入参信息
func (r *CainiaowaybilliigetAPIRequest) SetParamWaybillCloudPrintApplyNewRequest(_paramWaybillCloudPrintApplyNewRequest *WaybillCloudPrintApplyNewRequest) error {
	r._paramWaybillCloudPrintApplyNewRequest = _paramWaybillCloudPrintApplyNewRequest
	r.Set("param_waybill_cloud_print_apply_new_request", _paramWaybillCloudPrintApplyNewRequest)
	return nil
}

// GetParamWaybillCloudPrintApplyNewRequest ParamWaybillCloudPrintApplyNewRequest Getter
func (r CainiaowaybilliigetAPIRequest) GetParamWaybillCloudPrintApplyNewRequest() *WaybillCloudPrintApplyNewRequest {
	return r._paramWaybillCloudPrintApplyNewRequest
}
