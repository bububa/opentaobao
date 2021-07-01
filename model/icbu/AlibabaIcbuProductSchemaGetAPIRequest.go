package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductSchemaGetAPIRequest
（新）ICBU商品发布schema接口 API请求
alibaba.icbu.product.schema.get

获取ICBU商品发布的页面规则和填写字段，适用于新发商品 */
type AlibabaIcbuProductSchemaGetAPIRequest struct {
	model.Params
	// 商品规则渲染请求
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// New
