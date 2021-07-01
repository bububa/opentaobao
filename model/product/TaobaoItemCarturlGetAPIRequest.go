package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemCarturlGetAPIRequest
加购URL获取 API请求
taobao.item.carturl.get

获取加购URL，支持添加商品到购物车 */
type TaobaoItemCarturlGetAPIRequest struct {
	model.Params
	// 商品信息，格式为 商品ID_SKU ID_数量，多条记录以逗号(,)分割
	_itemIds []string
	// 回调地址，需要是EWS域名地址。可不填，默认到购物车页面
	_callbackUrl string
	// 商家Nick，优先使用user_id
	_userNick string
	// 商家ID
	_userId int64
	// 扩展属性，关注店铺的时候会传递下去，格式为K:V|K:V格式
	_extParams string
	// 端类型，默认是tb，可选tb,hm
	_type string
}

// New
