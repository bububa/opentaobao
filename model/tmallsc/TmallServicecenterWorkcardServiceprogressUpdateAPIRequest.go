package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardServiceprogressUpdateAPIRequest 回传工单服务进度 API请求
// tmall.servicecenter.workcard.serviceprogress.update
//
// 回传工单服务进度
type TmallServicecenterWorkcardServiceprogressUpdateAPIRequest struct {
	model.Params
	// 服务进度信息
	_updateServiceProgressRequest *UpdateServiceProgressRequest
}

// NewTmallServicecenterWorkcardServiceprogressUpdateRequest 初始化TmallServicecenterWorkcardServiceprogressUpdateAPIRequest对象
func NewTmallServicecenterWorkcardServiceprogressUpdateRequest() *TmallServicecenterWorkcardServiceprogressUpdateAPIRequest {
	return &TmallServicecenterWorkcardServiceprogressUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkcardServiceprogressUpdateAPIRequest) Reset() {
	r._updateServiceProgressRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardServiceprogressUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.serviceprogress.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardServiceprogressUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardServiceprogressUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateServiceProgressRequest is UpdateServiceProgressRequest Setter
// 服务进度信息
func (r *TmallServicecenterWorkcardServiceprogressUpdateAPIRequest) SetUpdateServiceProgressRequest(_updateServiceProgressRequest *UpdateServiceProgressRequest) error {
	r._updateServiceProgressRequest = _updateServiceProgressRequest
	r.Set("update_service_progress_request", _updateServiceProgressRequest)
	return nil
}

// GetUpdateServiceProgressRequest UpdateServiceProgressRequest Getter
func (r TmallServicecenterWorkcardServiceprogressUpdateAPIRequest) GetUpdateServiceProgressRequest() *UpdateServiceProgressRequest {
	return r._updateServiceProgressRequest
}

var poolTmallServicecenterWorkcardServiceprogressUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkcardServiceprogressUpdateRequest()
	},
}

// GetTmallServicecenterWorkcardServiceprogressUpdateRequest 从 sync.Pool 获取 TmallServicecenterWorkcardServiceprogressUpdateAPIRequest
func GetTmallServicecenterWorkcardServiceprogressUpdateAPIRequest() *TmallServicecenterWorkcardServiceprogressUpdateAPIRequest {
	return poolTmallServicecenterWorkcardServiceprogressUpdateAPIRequest.Get().(*TmallServicecenterWorkcardServiceprogressUpdateAPIRequest)
}

// ReleaseTmallServicecenterWorkcardServiceprogressUpdateAPIRequest 将 TmallServicecenterWorkcardServiceprogressUpdateAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkcardServiceprogressUpdateAPIRequest(v *TmallServicecenterWorkcardServiceprogressUpdateAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkcardServiceprogressUpdateAPIRequest.Put(v)
}
