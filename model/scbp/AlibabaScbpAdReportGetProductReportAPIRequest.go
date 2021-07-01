package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdReportGetProductReportAPIRequest
产品报告 API请求
alibaba.scbp.ad.report.get.product.report

产品报告 */
type AlibabaScbpAdReportGetProductReportAPIRequest struct {
	model.Params
	// 请求参数
	_productReportOperation *ProductReportOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
