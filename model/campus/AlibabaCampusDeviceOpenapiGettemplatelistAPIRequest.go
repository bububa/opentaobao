package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest 查询设备模板 API请求
// alibaba.campus.device.openapi.gettemplatelist
//
// 查询设备模板信息
type AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest struct {
	model.Params
	// 设备模板查询对象
	_query *TemplateApiQuery
}

// NewAlibabaCampusDeviceOpenapiGettemplatelistRequest 初始化AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest对象
func NewAlibabaCampusDeviceOpenapiGettemplatelistRequest() *AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest {
	return &AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.gettemplatelist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 设备模板查询对象
func (r *AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest) SetQuery(_query *TemplateApiQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest) GetQuery() *TemplateApiQuery {
	return r._query
}

var poolAlibabaCampusDeviceOpenapiGettemplatelistAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusDeviceOpenapiGettemplatelistRequest()
	},
}

// GetAlibabaCampusDeviceOpenapiGettemplatelistRequest 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest
func GetAlibabaCampusDeviceOpenapiGettemplatelistAPIRequest() *AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest {
	return poolAlibabaCampusDeviceOpenapiGettemplatelistAPIRequest.Get().(*AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest)
}

// ReleaseAlibabaCampusDeviceOpenapiGettemplatelistAPIRequest 将 AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiGettemplatelistAPIRequest(v *AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiGettemplatelistAPIRequest.Put(v)
}
