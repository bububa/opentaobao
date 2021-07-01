package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoShopSellerGetAPIRequest
卖家店铺基础信息查询 API请求
taobao.shop.seller.get

获取卖家店铺的基本信息 */
type TaobaoShopSellerGetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：Shop 结构中的所有字段；多个字段之间用逗号(,)分隔
	_fields string
}

// New
