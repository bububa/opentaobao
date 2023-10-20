package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthvaccinuserregisterremindAPIRequest isv到苗提醒 API请求
// alibaba.health.vaccin.user.register.remind
//
// isv到苗提醒
type AlibabahealthvaccinuserregisterremindAPIRequest struct {
	model.Params
	// 入参
	_isvVcAvailableRemindRequest *IsvVcAvailableRemindRequest
}

// NewAlibabahealthvaccinuserregisterremindRequest 初始化AlibabahealthvaccinuserregisterremindAPIRequest对象
func NewAlibabahealthvaccinuserregisterremindRequest() *AlibabahealthvaccinuserregisterremindAPIRequest {
	return &AlibabahealthvaccinuserregisterremindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahealthvaccinuserregisterremindAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.user.register.remind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahealthvaccinuserregisterremindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahealthvaccinuserregisterremindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvVcAvailableRemindRequest is IsvVcAvailableRemindRequest Setter
// 入参
func (r *AlibabahealthvaccinuserregisterremindAPIRequest) SetIsvVcAvailableRemindRequest(_isvVcAvailableRemindRequest *IsvVcAvailableRemindRequest) error {
	r._isvVcAvailableRemindRequest = _isvVcAvailableRemindRequest
	r.Set("isv_vc_available_remind_request", _isvVcAvailableRemindRequest)
	return nil
}

// GetIsvVcAvailableRemindRequest IsvVcAvailableRemindRequest Getter
func (r AlibabahealthvaccinuserregisterremindAPIRequest) GetIsvVcAvailableRemindRequest() *IsvVcAvailableRemindRequest {
	return r._isvVcAvailableRemindRequest
}
