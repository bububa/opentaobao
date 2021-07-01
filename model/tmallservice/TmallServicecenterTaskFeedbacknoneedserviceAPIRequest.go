package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterTaskFeedbacknoneedserviceAPIRequest
服务商反馈无需安装工单接口 API请求
tmall.servicecenter.task.feedbacknoneedservice

服务商反馈无需安装工单接口 */
type TmallServicecenterTaskFeedbacknoneedserviceAPIRequest struct {
	model.Params
	// 入参对象
	_param *SuspendServiceDo
}

// NewTmallServicecenterTaskFeedbacknoneedserviceRequest 初始化TmallServicecenterTaskFeedbacknoneedserviceAPIRequest对象
func NewTmallServicecenterTaskFeedbacknoneedserviceRequest() *TmallServicecenterTaskFeedbacknoneedserviceAPIRequest {
	return &TmallServicecenterTaskFeedbacknoneedserviceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterTaskFeedbacknoneedserviceAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.task.feedbacknoneedservice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterTaskFeedbacknoneedserviceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 入参对象
func (r *TmallServicecenterTaskFeedbacknoneedserviceAPIRequest) SetParam(_param *SuspendServiceDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r TmallServicecenterTaskFeedbacknoneedserviceAPIRequest) GetParam() *SuspendServiceDo {
	return r._param
}
