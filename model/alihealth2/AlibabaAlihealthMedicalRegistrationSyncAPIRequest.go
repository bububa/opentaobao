package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalregistrationsyncAPIRequest 阿里健康支付宝挂号记录回传接口 API请求
// alibaba.alihealth.medical.registration.sync
//
// 阿里健康支付宝挂号记录回传接口
type AlibabaalihealthmedicalregistrationsyncAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4top
}

// NewAlibabaalihealthmedicalregistrationsyncRequest 初始化AlibabaalihealthmedicalregistrationsyncAPIRequest对象
func NewAlibabaalihealthmedicalregistrationsyncRequest() *AlibabaalihealthmedicalregistrationsyncAPIRequest {
	return &AlibabaalihealthmedicalregistrationsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicalregistrationsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.registration.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicalregistrationsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicalregistrationsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSaveRequest is SaveRequest Setter
// 接口入参
func (r *AlibabaalihealthmedicalregistrationsyncAPIRequest) SetSaveRequest(_saveRequest *CommonRequest4top) error {
	r._saveRequest = _saveRequest
	r.Set("save_request", _saveRequest)
	return nil
}

// GetSaveRequest SaveRequest Getter
func (r AlibabaalihealthmedicalregistrationsyncAPIRequest) GetSaveRequest() *CommonRequest4top {
	return r._saveRequest
}
