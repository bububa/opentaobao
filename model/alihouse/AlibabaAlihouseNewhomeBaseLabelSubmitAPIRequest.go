package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) Reset() {
	r._labels = r._labels[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.base.label.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLabels is Labels Setter
// 标签列表
func (r *AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) SetLabels(_labels []BaseLabelDto) error {
	r._labels = _labels
	r.Set("labels", _labels)
	return nil
}

// GetLabels Labels Getter
func (r AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) GetLabels() []BaseLabelDto {
	return r._labels
}

var poolAlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeBaseLabelSubmitRequest()
	},
}

// GetAlibabaAlihouseNewhomeBaseLabelSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest
func GetAlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest() *AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest {
	return poolAlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest.Get().(*AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest 将 AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest(v *AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest.Put(v)
}
