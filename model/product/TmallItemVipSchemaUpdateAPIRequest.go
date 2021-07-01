package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemVipSchemaUpdateAPIRequest
大商家商品编辑接口 API请求
tmall.item.vip.schema.update

大商家编辑商品 */
type TmallItemVipSchemaUpdateAPIRequest struct {
	model.Params
	// 商品编辑的schema参数
	_schemaXmlFields string
	// 商品id
	_itemId int64
}

// New
