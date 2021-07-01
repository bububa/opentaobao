package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpEffectProductSingleGetAPIRequest
单个产品的报表 API请求
alibaba.scbp.effect.product.single.get

单个产品的报表 */
type AlibabaScbpEffectProductSingleGetAPIRequest struct {
	model.Params
	// ProductQuery
	_p4pProductReportQuery *ProductQuery
}

// New
