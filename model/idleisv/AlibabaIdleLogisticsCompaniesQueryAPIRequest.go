package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleLogisticsCompaniesQueryAPIRequest 快递公司列表查询 API请求
// alibaba.idle.logistics.companies.query
//
// 支持发货的快递公司列表查询
type AlibabaIdleLogisticsCompaniesQueryAPIRequest struct {
	model.Params
}

// NewAlibabaIdleLogisticsCompaniesQueryRequest 初始化AlibabaIdleLogisticsCompaniesQueryAPIRequest对象
func NewAlibabaIdleLogisticsCompaniesQueryRequest() *AlibabaIdleLogisticsCompaniesQueryAPIRequest {
	return &AlibabaIdleLogisticsCompaniesQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleLogisticsCompaniesQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.logistics.companies.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleLogisticsCompaniesQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleLogisticsCompaniesQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
