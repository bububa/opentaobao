package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectsubmitAPIRequest 提交楼盘信息 API请求
// alibaba.alihouse.newhome.project.submit
//
// 提交楼盘信息
type AlibabaalihousenewhomeprojectsubmitAPIRequest struct {
	model.Params
	// 楼盘对象
	_ebbasProjectDto *EbbasProjectDto
}

// NewAlibabaalihousenewhomeprojectsubmitRequest 初始化AlibabaalihousenewhomeprojectsubmitAPIRequest对象
func NewAlibabaalihousenewhomeprojectsubmitRequest() *AlibabaalihousenewhomeprojectsubmitAPIRequest {
	return &AlibabaalihousenewhomeprojectsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectsubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEbbasProjectDto is EbbasProjectDto Setter
// 楼盘对象
func (r *AlibabaalihousenewhomeprojectsubmitAPIRequest) SetEbbasProjectDto(_ebbasProjectDto *EbbasProjectDto) error {
	r._ebbasProjectDto = _ebbasProjectDto
	r.Set("ebbas_project_dto", _ebbasProjectDto)
	return nil
}

// GetEbbasProjectDto EbbasProjectDto Getter
func (r AlibabaalihousenewhomeprojectsubmitAPIRequest) GetEbbasProjectDto() *EbbasProjectDto {
	return r._ebbasProjectDto
}
