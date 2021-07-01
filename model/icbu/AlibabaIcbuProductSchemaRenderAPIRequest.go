package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductSchemaRenderAPIRequest
（新）获取商品信息 API请求
alibaba.icbu.product.schema.render

获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个商品编辑场景，不包括草稿。 */
type AlibabaIcbuProductSchemaRenderAPIRequest struct {
	model.Params
	// 商品规则渲染请求
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// New
