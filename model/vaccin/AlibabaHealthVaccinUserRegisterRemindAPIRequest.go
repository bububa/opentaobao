package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinUserRegisterRemindAPIRequest isv到苗提醒 API请求
// alibaba.health.vaccin.user.register.remind
//
// isv到苗提醒
type AlibabaHealthVaccinUserRegisterRemindAPIRequest struct {
	model.Params
	// 入参
	_isvVcAvailableRemindRequest *IsvVcAvailableRemindRequest
}

// NewAlibabaHealthVaccinUserRegisterRemindRequest 初始化AlibabaHealthVaccinUserRegisterRemindAPIRequest对象
func NewAlibabaHealthVaccinUserRegisterRemindRequest() *AlibabaHealthVaccinUserRegisterRemindAPIRequest {
	return &AlibabaHealthVaccinUserRegisterRemindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinUserRegisterRemindAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.user.register.remind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinUserRegisterRemindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHealthVaccinUserRegisterRemindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvVcAvailableRemindRequest is IsvVcAvailableRemindRequest Setter
// 入参
func (r *AlibabaHealthVaccinUserRegisterRemindAPIRequest) SetIsvVcAvailableRemindRequest(_isvVcAvailableRemindRequest *IsvVcAvailableRemindRequest) error {
	r._isvVcAvailableRemindRequest = _isvVcAvailableRemindRequest
	r.Set("isv_vc_available_remind_request", _isvVcAvailableRemindRequest)
	return nil
}

// GetIsvVcAvailableRemindRequest IsvVcAvailableRemindRequest Getter
func (r AlibabaHealthVaccinUserRegisterRemindAPIRequest) GetIsvVcAvailableRemindRequest() *IsvVcAvailableRemindRequest {
	return r._isvVcAvailableRemindRequest
}
