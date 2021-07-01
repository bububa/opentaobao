package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductSchemaAddDraftAPIRequest
（新）ICBU商品发布草稿 API请求
alibaba.icbu.product.schema.add.draft

提供发布ICBU商品草稿的入口 */
type AlibabaIcbuProductSchemaAddDraftAPIRequest struct {
	model.Params
	// 发布入参
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// New
