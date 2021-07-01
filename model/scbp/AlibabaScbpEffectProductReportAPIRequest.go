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

// New
