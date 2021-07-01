package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductSchemaUpdateAPIRequest
（新）商品发布增量更新接口 API请求
alibaba.icbu.product.schema.update

商品更新接口，方式为增量更新，只更新传入的字段 */
type AlibabaIcbuProductSchemaUpdateAPIRequest struct {
	model.Params
	// 发布入参
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// New
