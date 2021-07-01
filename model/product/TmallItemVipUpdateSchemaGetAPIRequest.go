package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemVipUpdateSchemaGetAPIRequest
vip商家编辑商品的规则获取接口 API请求
tmall.item.vip.update.schema.get

获取vip商家编辑商品的规则 */
type TmallItemVipUpdateSchemaGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// New
