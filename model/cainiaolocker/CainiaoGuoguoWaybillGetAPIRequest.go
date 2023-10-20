package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGuoguoWaybillGetAPIRequest 菜鸟裹裹商家寄件取号接口 API请求
// cainiao.guoguo.waybill.get
//
// 菜鸟裹裹商家寄件取号接口
type CainiaoGuoguoWaybillGetAPIRequest struct {
	model.Params
	// 入参信息
	_paramWaybillCloudPrintApplyNewRequest *WaybillCloudPrintApplyNewRequest
}

// NewCainiaoGuoguoWaybillGetRequest 初始化CainiaoGuoguoWaybillGetAPIRequest对象
func NewCainiaoGuoguoWaybillGetRequest() *CainiaoGuoguoWaybillGetAPIRequest {
	return &CainiaoGuoguoWaybillGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGuoguoWaybillGetAPIRequest) GetApiMethodName() string {
	return "cainiao.guoguo.waybill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGuoguoWaybillGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGuoguoWaybillGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamWaybillCloudPrintApplyNewRequest is ParamWaybillCloudPrintApplyNewRequest Setter
// 入参信息
func (r *CainiaoGuoguoWaybillGetAPIRequest) SetParamWaybillCloudPrintApplyNewRequest(_paramWaybillCloudPrintApplyNewRequest *WaybillCloudPrintApplyNewRequest) error {
	r._paramWaybillCloudPrintApplyNewRequest = _paramWaybillCloudPrintApplyNewRequest
	r.Set("param_waybill_cloud_print_apply_new_request", _paramWaybillCloudPrintApplyNewRequest)
	return nil
}

// GetParamWaybillCloudPrintApplyNewRequest ParamWaybillCloudPrintApplyNewRequest Getter
func (r CainiaoGuoguoWaybillGetAPIRequest) GetParamWaybillCloudPrintApplyNewRequest() *WaybillCloudPrintApplyNewRequest {
	return r._paramWaybillCloudPrintApplyNewRequest
}
