package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest 提交标签库 API请求
// alibaba.alihouse.newhome.base.label.submit
//
// 提交标签库
type AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest struct {
	model.Params
	// 标签列表
	_labels []BaseLabelDto
}

// NewAlibabaAlihouseNewhomeBaseLabelSubmitRequest 初始化AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeBaseLabelSubmitRequest() *AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.base.label.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Labels Setter
// 标签列表
func (r *AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) SetLabels(_labels []BaseLabelDto) error {
	r._labels = _labels
	r.Set("labels", _labels)
	return nil
}

// Get Labels Getter
func (r AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) GetLabels() []BaseLabelDto {
	return r._labels
}
