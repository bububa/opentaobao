package waybill

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiGetAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamWaybillCloudPrintApplyNewRequest Setter
// 入参信息
func (r *CainiaoWaybillIiGetAPIRequest) SetParamWaybillCloudPrintApplyNewRequest(_paramWaybillCloudPrintApplyNewRequest *WaybillCloudPrintApplyNewRequest) error {
	r._paramWaybillCloudPrintApplyNewRequest = _paramWaybillCloudPrintApplyNewRequest
	r.Set("param_waybill_cloud_print_apply_new_request", _paramWaybillCloudPrintApplyNewRequest)
	return nil
}

// Get ParamWaybillCloudPrintApplyNewRequest Getter
func (r CainiaoWaybillIiGetAPIRequest) GetParamWaybillCloudPrintApplyNewRequest() *WaybillCloudPrintApplyNewRequest {
	return r._paramWaybillCloudPrintApplyNewRequest
}
