package alidoc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthnrrxprescriptiongetAPIRequest 搜索处方详情 API请求
// alibaba.alihealth.nr.rx.prescription.get
//
// 获取互联网医院处方详情
type AlibabaalihealthnrrxprescriptiongetAPIRequest struct {
	model.Params
	// 查询参数
	_query *PrescriptionQueryDto
}

// NewAlibabaalihealthnrrxprescriptiongetRequest 初始化AlibabaalihealthnrrxprescriptiongetAPIRequest对象
func NewAlibabaalihealthnrrxprescriptiongetRequest() *AlibabaalihealthnrrxprescriptiongetAPIRequest {
	return &AlibabaalihealthnrrxprescriptiongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthnrrxprescriptiongetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nr.rx.prescription.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthnrrxprescriptiongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthnrrxprescriptiongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询参数
func (r *AlibabaalihealthnrrxprescriptiongetAPIRequest) SetQuery(_query *PrescriptionQueryDto) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaalihealthnrrxprescriptiongetAPIRequest) GetQuery() *PrescriptionQueryDto {
	return r._query
}
