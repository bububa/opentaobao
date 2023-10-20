package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlelogisticscompaniesqueryAPIRequest 快递公司列表查询 API请求
// alibaba.idle.logistics.companies.query
//
// 支持发货的快递公司列表查询
type AlibabaidlelogisticscompaniesqueryAPIRequest struct {
	model.Params
}

// NewAlibabaidlelogisticscompaniesqueryRequest 初始化AlibabaidlelogisticscompaniesqueryAPIRequest对象
func NewAlibabaidlelogisticscompaniesqueryRequest() *AlibabaidlelogisticscompaniesqueryAPIRequest {
	return &AlibabaidlelogisticscompaniesqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlelogisticscompaniesqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.logistics.companies.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlelogisticscompaniesqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlelogisticscompaniesqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
