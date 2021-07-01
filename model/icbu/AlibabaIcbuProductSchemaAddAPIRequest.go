package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductSchemaAddAPIRequest
（新）商品发布新接口 API请求
alibaba.icbu.product.schema.add

提供发布ICBU商品的入口 */
type AlibabaIcbuProductSchemaAddAPIRequest struct {
	model.Params
	// 发布入参
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// New
