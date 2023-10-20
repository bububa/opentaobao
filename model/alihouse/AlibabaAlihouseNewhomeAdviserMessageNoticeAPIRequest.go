package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeadvisermessagenoticeAPIRequest 催促小B发送短信 API请求
// alibaba.alihouse.newhome.adviser.message.notice
//
// 催促小B发送短信
type AlibabaalihousenewhomeadvisermessagenoticeAPIRequest struct {
	model.Params
	// 置业顾问/经纪人对象
	_projectAdviserDto *ProjectAdviserDto
}

// NewAlibabaalihousenewhomeadvisermessagenoticeRequest 初始化AlibabaalihousenewhomeadvisermessagenoticeAPIRequest对象
func NewAlibabaalihousenewhomeadvisermessagenoticeRequest() *AlibabaalihousenewhomeadvisermessagenoticeAPIRequest {
	return &AlibabaalihousenewhomeadvisermessagenoticeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeadvisermessagenoticeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.adviser.message.notice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeadvisermessagenoticeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeadvisermessagenoticeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectAdviserDto is ProjectAdviserDto Setter
// 置业顾问/经纪人对象
func (r *AlibabaalihousenewhomeadvisermessagenoticeAPIRequest) SetProjectAdviserDto(_projectAdviserDto *ProjectAdviserDto) error {
	r._projectAdviserDto = _projectAdviserDto
	r.Set("project_adviser_dto", _projectAdviserDto)
	return nil
}

// GetProjectAdviserDto ProjectAdviserDto Getter
func (r AlibabaalihousenewhomeadvisermessagenoticeAPIRequest) GetProjectAdviserDto() *ProjectAdviserDto {
	return r._projectAdviserDto
}
