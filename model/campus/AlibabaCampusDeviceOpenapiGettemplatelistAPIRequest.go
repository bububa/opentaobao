package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusdeviceopenapigettemplatelistAPIRequest 查询设备模板 API请求
// alibaba.campus.device.openapi.gettemplatelist
//
// 查询设备模板信息
type AlibabacampusdeviceopenapigettemplatelistAPIRequest struct {
	model.Params
	// 设备模板查询对象
	_query *TemplateApiQuery
}

// NewAlibabacampusdeviceopenapigettemplatelistRequest 初始化AlibabacampusdeviceopenapigettemplatelistAPIRequest对象
func NewAlibabacampusdeviceopenapigettemplatelistRequest() *AlibabacampusdeviceopenapigettemplatelistAPIRequest {
	return &AlibabacampusdeviceopenapigettemplatelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusdeviceopenapigettemplatelistAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.gettemplatelist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusdeviceopenapigettemplatelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusdeviceopenapigettemplatelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 设备模板查询对象
func (r *AlibabacampusdeviceopenapigettemplatelistAPIRequest) SetQuery(_query *TemplateApiQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabacampusdeviceopenapigettemplatelistAPIRequest) GetQuery() *TemplateApiQuery {
	return r._query
}
