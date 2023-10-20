package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterTaskFeedbacknoneedserviceAPIRequest 服务商反馈无需安装工单接口 API请求
// tmall.servicecenter.task.feedbacknoneedservice
//
// 服务商反馈无需安装工单接口
type TmallServicecenterTaskFeedbacknoneedserviceAPIRequest struct {
	model.Params
	// 入参对象
	_param *SuspendServiceDo
}

// NewTmallServicecenterTaskFeedbacknoneedserviceRequest 初始化TmallServicecenterTaskFeedbacknoneedserviceAPIRequest对象
func NewTmallServicecenterTaskFeedbacknoneedserviceRequest() *TmallServicecenterTaskFeedbacknoneedserviceAPIRequest {
	return &TmallServicecenterTaskFeedbacknoneedserviceAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterTaskFeedbacknoneedserviceAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterTaskFeedbacknoneedserviceAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.task.feedbacknoneedservice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterTaskFeedbacknoneedserviceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterTaskFeedbacknoneedserviceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参对象
func (r *TmallServicecenterTaskFeedbacknoneedserviceAPIRequest) SetParam(_param *SuspendServiceDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TmallServicecenterTaskFeedbacknoneedserviceAPIRequest) GetParam() *SuspendServiceDo {
	return r._param
}

var poolTmallServicecenterTaskFeedbacknoneedserviceAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterTaskFeedbacknoneedserviceRequest()
	},
}

// GetTmallServicecenterTaskFeedbacknoneedserviceRequest 从 sync.Pool 获取 TmallServicecenterTaskFeedbacknoneedserviceAPIRequest
func GetTmallServicecenterTaskFeedbacknoneedserviceAPIRequest() *TmallServicecenterTaskFeedbacknoneedserviceAPIRequest {
	return poolTmallServicecenterTaskFeedbacknoneedserviceAPIRequest.Get().(*TmallServicecenterTaskFeedbacknoneedserviceAPIRequest)
}

// ReleaseTmallServicecenterTaskFeedbacknoneedserviceAPIRequest 将 TmallServicecenterTaskFeedbacknoneedserviceAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterTaskFeedbacknoneedserviceAPIRequest(v *TmallServicecenterTaskFeedbacknoneedserviceAPIRequest) {
	v.Reset()
	poolTmallServicecenterTaskFeedbacknoneedserviceAPIRequest.Put(v)
}
