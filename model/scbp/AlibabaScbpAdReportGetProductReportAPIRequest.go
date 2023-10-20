package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadreportgetproductreportAPIRequest 产品报告 API请求
// alibaba.scbp.ad.report.get.product.report
//
// 产品报告
type AlibabascbpadreportgetproductreportAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求参数
	_productReportOperation *ProductReportOperationDto
}

// NewAlibabascbpadreportgetproductreportRequest 初始化AlibabascbpadreportgetproductreportAPIRequest对象
func NewAlibabascbpadreportgetproductreportRequest() *AlibabascbpadreportgetproductreportAPIRequest {
	return &AlibabascbpadreportgetproductreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadreportgetproductreportAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.get.product.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadreportgetproductreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadreportgetproductreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadreportgetproductreportAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadreportgetproductreportAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetProductReportOperation is ProductReportOperation Setter
// 请求参数
func (r *AlibabascbpadreportgetproductreportAPIRequest) SetProductReportOperation(_productReportOperation *ProductReportOperationDto) error {
	r._productReportOperation = _productReportOperation
	r.Set("product_report_operation", _productReportOperation)
	return nil
}

// GetProductReportOperation ProductReportOperation Getter
func (r AlibabascbpadreportgetproductreportAPIRequest) GetProductReportOperation() *ProductReportOperationDto {
	return r._productReportOperation
}
