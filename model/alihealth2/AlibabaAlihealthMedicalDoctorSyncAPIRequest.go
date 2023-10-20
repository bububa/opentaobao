package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicaldoctorsyncAPIRequest 阿里健康预约挂号医生同步接口 API请求
// alibaba.alihealth.medical.doctor.sync
//
// 阿里健康预约挂号医生同步接口
type AlibabaalihealthmedicaldoctorsyncAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4top
}

// NewAlibabaalihealthmedicaldoctorsyncRequest 初始化AlibabaalihealthmedicaldoctorsyncAPIRequest对象
func NewAlibabaalihealthmedicaldoctorsyncRequest() *AlibabaalihealthmedicaldoctorsyncAPIRequest {
	return &AlibabaalihealthmedicaldoctorsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicaldoctorsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.doctor.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicaldoctorsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicaldoctorsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSaveRequest is SaveRequest Setter
// 接口入参
func (r *AlibabaalihealthmedicaldoctorsyncAPIRequest) SetSaveRequest(_saveRequest *CommonRequest4top) error {
	r._saveRequest = _saveRequest
	r.Set("save_request", _saveRequest)
	return nil
}

// GetSaveRequest SaveRequest Getter
func (r AlibabaalihealthmedicaldoctorsyncAPIRequest) GetSaveRequest() *CommonRequest4top {
	return r._saveRequest
}
