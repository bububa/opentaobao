package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectSubmitAPIRequest 提交楼盘信息 API请求
// alibaba.alihouse.newhome.project.submit
//
// 提交楼盘信息
type AlibabaAlihouseNewhomeProjectSubmitAPIRequest struct {
	model.Params
	// 楼盘对象
	_ebbasProjectDto *EbbasProjectDto
}

// NewAlibabaAlihouseNewhomeProjectSubmitRequest 初始化AlibabaAlihouseNewhomeProjectSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectSubmitRequest() *AlibabaAlihouseNewhomeProjectSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeProjectSubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectSubmitAPIRequest) Reset() {
	r._ebbasProjectDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEbbasProjectDto is EbbasProjectDto Setter
// 楼盘对象
func (r *AlibabaAlihouseNewhomeProjectSubmitAPIRequest) SetEbbasProjectDto(_ebbasProjectDto *EbbasProjectDto) error {
	r._ebbasProjectDto = _ebbasProjectDto
	r.Set("ebbas_project_dto", _ebbasProjectDto)
	return nil
}

// GetEbbasProjectDto EbbasProjectDto Getter
func (r AlibabaAlihouseNewhomeProjectSubmitAPIRequest) GetEbbasProjectDto() *EbbasProjectDto {
	return r._ebbasProjectDto
}

var poolAlibabaAlihouseNewhomeProjectSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectSubmitRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectSubmitAPIRequest
func GetAlibabaAlihouseNewhomeProjectSubmitAPIRequest() *AlibabaAlihouseNewhomeProjectSubmitAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectSubmitAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectSubmitAPIRequest 将 AlibabaAlihouseNewhomeProjectSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectSubmitAPIRequest(v *AlibabaAlihouseNewhomeProjectSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectSubmitAPIRequest.Put(v)
}
