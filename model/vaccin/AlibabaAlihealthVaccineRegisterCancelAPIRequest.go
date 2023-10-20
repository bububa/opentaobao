package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthVaccineRegisterCancelAPIRequest 取消登记 API请求
// alibaba.alihealth.vaccine.register.cancel
//
// 取消登记
type AlibabaAlihealthVaccineRegisterCancelAPIRequest struct {
	model.Params
	// 无
	_topRequest *CancelVcRegisterRequest
}

// NewAlibabaAlihealthVaccineRegisterCancelRequest 初始化AlibabaAlihealthVaccineRegisterCancelAPIRequest对象
func NewAlibabaAlihealthVaccineRegisterCancelRequest() *AlibabaAlihealthVaccineRegisterCancelAPIRequest {
	return &AlibabaAlihealthVaccineRegisterCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthVaccineRegisterCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.vaccine.register.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthVaccineRegisterCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthVaccineRegisterCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopRequest is TopRequest Setter
// 无
func (r *AlibabaAlihealthVaccineRegisterCancelAPIRequest) SetTopRequest(_topRequest *CancelVcRegisterRequest) error {
	r._topRequest = _topRequest
	r.Set("top_request", _topRequest)
	return nil
}

// GetTopRequest TopRequest Getter
func (r AlibabaAlihealthVaccineRegisterCancelAPIRequest) GetTopRequest() *CancelVcRegisterRequest {
	return r._topRequest
}
