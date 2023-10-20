package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest 置业顾问批量绑定楼盘 API请求
// alibaba.alihouse.newhome.project.adviser.bind
//
// 置业顾问批量绑定楼盘
type AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest struct {
	model.Params
	// 对象
	_projectAdviserBindDto *ProjectAdviserBindDto
	// 状态：0无效，1有效
	_status int64
}

// NewAlibabaAlihouseNewhomeProjectAdviserBindRequest 初始化AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectAdviserBindRequest() *AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest {
	return &AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) Reset() {
	r._projectAdviserBindDto = nil
	r._status = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.adviser.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectAdviserBindDto is ProjectAdviserBindDto Setter
// 对象
func (r *AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) SetProjectAdviserBindDto(_projectAdviserBindDto *ProjectAdviserBindDto) error {
	r._projectAdviserBindDto = _projectAdviserBindDto
	r.Set("project_adviser_bind_dto", _projectAdviserBindDto)
	return nil
}

// GetProjectAdviserBindDto ProjectAdviserBindDto Getter
func (r AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) GetProjectAdviserBindDto() *ProjectAdviserBindDto {
	return r._projectAdviserBindDto
}

// SetStatus is Status Setter
// 状态：0无效，1有效
func (r *AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) GetStatus() int64 {
	return r._status
}

var poolAlibabaAlihouseNewhomeProjectAdviserBindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectAdviserBindRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectAdviserBindRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest
func GetAlibabaAlihouseNewhomeProjectAdviserBindAPIRequest() *AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectAdviserBindAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectAdviserBindAPIRequest 将 AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectAdviserBindAPIRequest(v *AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectAdviserBindAPIRequest.Put(v)
}
