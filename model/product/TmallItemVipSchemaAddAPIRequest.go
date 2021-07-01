package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemVipSchemaAddAPIRequest
大商家商品发布接口 API请求
tmall.item.vip.schema.add

大商家商品发布接口 */
type TmallItemVipSchemaAddAPIRequest struct {
	model.Params
	// 商品发布schema参数
	_schemaXmlFields string
}

// New
