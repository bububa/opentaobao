package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardserviceprogressupdateAPIRequest 回传工单服务进度 API请求
// tmall.servicecenter.workcard.serviceprogress.update
//
// 回传工单服务进度
type TmallservicecenterworkcardserviceprogressupdateAPIRequest struct {
	model.Params
	// 服务进度信息
	_updateServiceProgressRequest *UpdateServiceProgressRequest
}

// NewTmallservicecenterworkcardserviceprogressupdateRequest 初始化TmallservicecenterworkcardserviceprogressupdateAPIRequest对象
func NewTmallservicecenterworkcardserviceprogressupdateRequest() *TmallservicecenterworkcardserviceprogressupdateAPIRequest {
	return &TmallservicecenterworkcardserviceprogressupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardserviceprogressupdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.serviceprogress.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardserviceprogressupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardserviceprogressupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateServiceProgressRequest is UpdateServiceProgressRequest Setter
// 服务进度信息
func (r *TmallservicecenterworkcardserviceprogressupdateAPIRequest) SetUpdateServiceProgressRequest(_updateServiceProgressRequest *UpdateServiceProgressRequest) error {
	r._updateServiceProgressRequest = _updateServiceProgressRequest
	r.Set("update_service_progress_request", _updateServiceProgressRequest)
	return nil
}

// GetUpdateServiceProgressRequest UpdateServiceProgressRequest Getter
func (r TmallservicecenterworkcardserviceprogressupdateAPIRequest) GetUpdateServiceProgressRequest() *UpdateServiceProgressRequest {
	return r._updateServiceProgressRequest
}
