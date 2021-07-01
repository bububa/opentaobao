package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest
查询设备模板 API请求
alibaba.campus.device.openapi.gettemplatelist

查询设备模板信息 */
type AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest struct {
	model.Params
	// 设备模板查询对象
	_query *TemplateApiQuery
}

// NewAlibabaCampusDeviceOpenapiGettemplatelistRequest 初始化AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest对象
func NewAlibabaCampusDeviceOpenapiGettemplatelistRequest() *AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest {
	return &AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.gettemplatelist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Query Setter
// 设备模板查询对象
func (r *AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest) SetQuery(_query *TemplateApiQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// Get Query Getter
func (r AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest) GetQuery() *TemplateApiQuery {
	return r._query
}
