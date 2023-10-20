package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectadvisersubmitAPIRequest 提交楼盘顾问 API请求
// alibaba.alihouse.newhome.project.adviser.submit
//
// 提交楼盘顾问
type AlibabaalihousenewhomeprojectadvisersubmitAPIRequest struct {
	model.Params
	// 顾问列表
	_advisers []ProjectAdviserDto
}

// NewAlibabaalihousenewhomeprojectadvisersubmitRequest 初始化AlibabaalihousenewhomeprojectadvisersubmitAPIRequest对象
func NewAlibabaalihousenewhomeprojectadvisersubmitRequest() *AlibabaalihousenewhomeprojectadvisersubmitAPIRequest {
	return &AlibabaalihousenewhomeprojectadvisersubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectadvisersubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.adviser.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectadvisersubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectadvisersubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdvisers is Advisers Setter
// 顾问列表
func (r *AlibabaalihousenewhomeprojectadvisersubmitAPIRequest) SetAdvisers(_advisers []ProjectAdviserDto) error {
	r._advisers = _advisers
	r.Set("advisers", _advisers)
	return nil
}

// GetAdvisers Advisers Getter
func (r AlibabaalihousenewhomeprojectadvisersubmitAPIRequest) GetAdvisers() []ProjectAdviserDto {
	return r._advisers
}
