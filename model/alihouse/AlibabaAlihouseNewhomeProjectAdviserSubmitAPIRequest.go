package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest 提交楼盘顾问 API请求
// alibaba.alihouse.newhome.project.adviser.submit
//
// 提交楼盘顾问
type AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest struct {
	model.Params
	// 顾问列表
	_advisers []ProjectAdviserDto
}

// NewAlibabaAlihouseNewhomeProjectAdviserSubmitRequest 初始化AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectAdviserSubmitRequest() *AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest) Reset() {
	r._advisers = r._advisers[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.adviser.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdvisers is Advisers Setter
// 顾问列表
func (r *AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest) SetAdvisers(_advisers []ProjectAdviserDto) error {
	r._advisers = _advisers
	r.Set("advisers", _advisers)
	return nil
}

// GetAdvisers Advisers Getter
func (r AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest) GetAdvisers() []ProjectAdviserDto {
	return r._advisers
}

var poolAlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectAdviserSubmitRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectAdviserSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest
func GetAlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest() *AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest 将 AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest(v *AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest.Put(v)
}
