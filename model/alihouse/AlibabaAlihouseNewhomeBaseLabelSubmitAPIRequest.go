package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomebaselabelsubmitAPIRequest 提交标签库 API请求
// alibaba.alihouse.newhome.base.label.submit
//
// 提交标签库
type AlibabaalihousenewhomebaselabelsubmitAPIRequest struct {
	model.Params
	// 标签列表
	_labels []BaseLabelDto
}

// NewAlibabaalihousenewhomebaselabelsubmitRequest 初始化AlibabaalihousenewhomebaselabelsubmitAPIRequest对象
func NewAlibabaalihousenewhomebaselabelsubmitRequest() *AlibabaalihousenewhomebaselabelsubmitAPIRequest {
	return &AlibabaalihousenewhomebaselabelsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomebaselabelsubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.base.label.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomebaselabelsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomebaselabelsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLabels is Labels Setter
// 标签列表
func (r *AlibabaalihousenewhomebaselabelsubmitAPIRequest) SetLabels(_labels []BaseLabelDto) error {
	r._labels = _labels
	r.Set("labels", _labels)
	return nil
}

// GetLabels Labels Getter
func (r AlibabaalihousenewhomebaselabelsubmitAPIRequest) GetLabels() []BaseLabelDto {
	return r._labels
}
