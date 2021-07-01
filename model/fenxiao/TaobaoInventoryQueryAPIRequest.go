package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryQueryAPIRequest
查询商品库存信息 API请求
taobao.inventory.query

建议使用新接口：tmall.inventory.query.forstore ，新ISV不推荐使用。
商家查询商品总体库存信息 */
type TaobaoInventoryQueryAPIRequest struct {
	model.Params
	// 后端商品ID 列表，控制到50个
	_scItemIds string
	// 后端商品的商家编码列表，控制到50个
	_scItemCodes string
	// 卖家昵称
	_sellerNick string
	// 仓库列表:GLY001^GLY002
	_storeCodes string
}

// New
