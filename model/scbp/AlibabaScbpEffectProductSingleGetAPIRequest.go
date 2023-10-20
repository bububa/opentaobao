package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpeffectproductsinglegetAPIRequest 单个产品的报表 API请求
// alibaba.scbp.effect.product.single.get
//
// 单个产品的报表
type AlibabascbpeffectproductsinglegetAPIRequest struct {
	model.Params
	// ProductQuery
	_p4pProductReportQuery *ProductQuery
}

// NewAlibabascbpeffectproductsinglegetRequest 初始化AlibabascbpeffectproductsinglegetAPIRequest对象
func NewAlibabascbpeffectproductsinglegetRequest() *AlibabascbpeffectproductsinglegetAPIRequest {
	return &AlibabascbpeffectproductsinglegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpeffectproductsinglegetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.effect.product.single.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpeffectproductsinglegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpeffectproductsinglegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetP4pProductReportQuery is P4pProductReportQuery Setter
// ProductQuery
func (r *AlibabascbpeffectproductsinglegetAPIRequest) SetP4pProductReportQuery(_p4pProductReportQuery *ProductQuery) error {
	r._p4pProductReportQuery = _p4pProductReportQuery
	r.Set("p4p_product_report_query", _p4pProductReportQuery)
	return nil
}

// GetP4pProductReportQuery P4pProductReportQuery Getter
func (r AlibabascbpeffectproductsinglegetAPIRequest) GetP4pProductReportQuery() *ProductQuery {
	return r._p4pProductReportQuery
}
