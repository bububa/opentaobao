package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemItemPublishAPIRequest
全渠道门店商品轻发布 API请求
taobao.omniitem.item.publish

全渠道门店商品轻发布 */
type TaobaoOmniitemItemPublishAPIRequest struct {
	model.Params
	// 发布商品信息
	_lightPublishInfo *ItemLightPublishDto
	// 在全域商品或是门店商品中校验码是否重复，可选值对应为ALL或者STORE
	_operateType string
}

// New
