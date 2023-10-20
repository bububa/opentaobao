package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoguoguowaybillgetAPIRequest 菜鸟裹裹商家寄件取号接口 API请求
// cainiao.guoguo.waybill.get
//
// 菜鸟裹裹商家寄件取号接口
type CainiaoguoguowaybillgetAPIRequest struct {
	model.Params
	// 入参信息
	_paramWaybillCloudPrintApplyNewRequest *WaybillCloudPrintApplyNewRequest
}

// NewCainiaoguoguowaybillgetRequest 初始化CainiaoguoguowaybillgetAPIRequest对象
func NewCainiaoguoguowaybillgetRequest() *CainiaoguoguowaybillgetAPIRequest {
	return &CainiaoguoguowaybillgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoguoguowaybillgetAPIRequest) GetApiMethodName() string {
	return "cainiao.guoguo.waybill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoguoguowaybillgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoguoguowaybillgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamWaybillCloudPrintApplyNewRequest is ParamWaybillCloudPrintApplyNewRequest Setter
// 入参信息
func (r *CainiaoguoguowaybillgetAPIRequest) SetParamWaybillCloudPrintApplyNewRequest(_paramWaybillCloudPrintApplyNewRequest *WaybillCloudPrintApplyNewRequest) error {
	r._paramWaybillCloudPrintApplyNewRequest = _paramWaybillCloudPrintApplyNewRequest
	r.Set("param_waybill_cloud_print_apply_new_request", _paramWaybillCloudPrintApplyNewRequest)
	return nil
}

// GetParamWaybillCloudPrintApplyNewRequest ParamWaybillCloudPrintApplyNewRequest Getter
func (r CainiaoguoguowaybillgetAPIRequest) GetParamWaybillCloudPrintApplyNewRequest() *WaybillCloudPrintApplyNewRequest {
	return r._paramWaybillCloudPrintApplyNewRequest
}
