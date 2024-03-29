package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportGetProductReportAPIRequest 产品报告 API请求
// alibaba.scbp.ad.report.get.product.report
//
// 产品报告
type AlibabaScbpAdReportGetProductReportAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求参数
	_productReportOperation *ProductReportOperationDto
}

// NewAlibabaScbpAdReportGetProductReportRequest 初始化AlibabaScbpAdReportGetProductReportAPIRequest对象
func NewAlibabaScbpAdReportGetProductReportRequest() *AlibabaScbpAdReportGetProductReportAPIRequest {
	return &AlibabaScbpAdReportGetProductReportAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdReportGetProductReportAPIRequest) Reset() {
	r._topContext = nil
	r._productReportOperation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportGetProductReportAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.get.product.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportGetProductReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdReportGetProductReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportGetProductReportAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdReportGetProductReportAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetProductReportOperation is ProductReportOperation Setter
// 请求参数
func (r *AlibabaScbpAdReportGetProductReportAPIRequest) SetProductReportOperation(_productReportOperation *ProductReportOperationDto) error {
	r._productReportOperation = _productReportOperation
	r.Set("product_report_operation", _productReportOperation)
	return nil
}

// GetProductReportOperation ProductReportOperation Getter
func (r AlibabaScbpAdReportGetProductReportAPIRequest) GetProductReportOperation() *ProductReportOperationDto {
	return r._productReportOperation
}

var poolAlibabaScbpAdReportGetProductReportAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdReportGetProductReportRequest()
	},
}

// GetAlibabaScbpAdReportGetProductReportRequest 从 sync.Pool 获取 AlibabaScbpAdReportGetProductReportAPIRequest
func GetAlibabaScbpAdReportGetProductReportAPIRequest() *AlibabaScbpAdReportGetProductReportAPIRequest {
	return poolAlibabaScbpAdReportGetProductReportAPIRequest.Get().(*AlibabaScbpAdReportGetProductReportAPIRequest)
}

// ReleaseAlibabaScbpAdReportGetProductReportAPIRequest 将 AlibabaScbpAdReportGetProductReportAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdReportGetProductReportAPIRequest(v *AlibabaScbpAdReportGetProductReportAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdReportGetProductReportAPIRequest.Put(v)
}
