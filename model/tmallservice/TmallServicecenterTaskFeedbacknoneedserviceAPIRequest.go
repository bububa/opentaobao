package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecentertaskfeedbacknoneedserviceAPIRequest 服务商反馈无需安装工单接口 API请求
// tmall.servicecenter.task.feedbacknoneedservice
//
// 服务商反馈无需安装工单接口
type TmallservicecentertaskfeedbacknoneedserviceAPIRequest struct {
	model.Params
	// 入参对象
	_param *SuspendServiceDo
}

// NewTmallservicecentertaskfeedbacknoneedserviceRequest 初始化TmallservicecentertaskfeedbacknoneedserviceAPIRequest对象
func NewTmallservicecentertaskfeedbacknoneedserviceRequest() *TmallservicecentertaskfeedbacknoneedserviceAPIRequest {
	return &TmallservicecentertaskfeedbacknoneedserviceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecentertaskfeedbacknoneedserviceAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.task.feedbacknoneedservice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecentertaskfeedbacknoneedserviceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecentertaskfeedbacknoneedserviceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参对象
func (r *TmallservicecentertaskfeedbacknoneedserviceAPIRequest) SetParam(_param *SuspendServiceDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TmallservicecentertaskfeedbacknoneedserviceAPIRequest) GetParam() *SuspendServiceDo {
	return r._param
}
