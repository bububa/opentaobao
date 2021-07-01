package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpEffectProductReportAPIRequest
所有产品报表 API请求
alibaba.scbp.effect.product.report

所有产品报表 */
type AlibabaScbpEffectProductReportAPIRequest struct {
	model.Params
	// ProductQuery
	_p4pProductReportQuery *ProductQuery
}

// NewAlibabaScbpEffectProductReportRequest 初始化AlibabaScbpEffectProductReportAPIRequest对象
func NewAlibabaScbpEffectProductReportRequest() *AlibabaScbpEffectProductReportAPIRequest {
	return &AlibabaScbpEffectProductReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectProductReportAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.effect.product.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectProductReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is P4pProductReportQuery Setter
// ProductQuery
func (r *AlibabaScbpEffectProductReportAPIRequest) SetP4pProductReportQuery(_p4pProductReportQuery *ProductQuery) error {
	r._p4pProductReportQuery = _p4pProductReportQuery
	r.Set("p4p_product_report_query", _p4pProductReportQuery)
	return nil
}

// Get P4pProductReportQuery Getter
func (r AlibabaScbpEffectProductReportAPIRequest) GetP4pProductReportQuery() *ProductQuery {
	return r._p4pProductReportQuery
}
