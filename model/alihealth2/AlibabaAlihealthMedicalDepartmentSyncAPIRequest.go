package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicaldepartmentsyncAPIRequest 阿里健康预约挂号科室同步接口 API请求
// alibaba.alihealth.medical.department.sync
//
// 阿里健康预约挂号科室同步接口
type AlibabaalihealthmedicaldepartmentsyncAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4top
}

// NewAlibabaalihealthmedicaldepartmentsyncRequest 初始化AlibabaalihealthmedicaldepartmentsyncAPIRequest对象
func NewAlibabaalihealthmedicaldepartmentsyncRequest() *AlibabaalihealthmedicaldepartmentsyncAPIRequest {
	return &AlibabaalihealthmedicaldepartmentsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicaldepartmentsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.department.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicaldepartmentsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicaldepartmentsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSaveRequest is SaveRequest Setter
// 接口入参
func (r *AlibabaalihealthmedicaldepartmentsyncAPIRequest) SetSaveRequest(_saveRequest *CommonRequest4top) error {
	r._saveRequest = _saveRequest
	r.Set("save_request", _saveRequest)
	return nil
}

// GetSaveRequest SaveRequest Getter
func (r AlibabaalihealthmedicaldepartmentsyncAPIRequest) GetSaveRequest() *CommonRequest4top {
	return r._saveRequest
}
