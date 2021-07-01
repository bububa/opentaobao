package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductSchemaRenderDraftAPIRequest
（新）渲染草稿商品数据 API请求
alibaba.icbu.product.schema.render.draft

获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个草稿商品编辑场景， */
type AlibabaIcbuProductSchemaRenderDraftAPIRequest struct {
	model.Params
	// 商品规则渲染请求
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// New
