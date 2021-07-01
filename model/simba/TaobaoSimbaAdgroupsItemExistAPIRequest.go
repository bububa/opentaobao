package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaAdgroupsItemExistAPIRequest
商品是否推广 API请求
taobao.simba.adgroups.item.exist

判断在一个推广计划中是否已经推广了一个商品 */
type TaobaoSimbaAdgroupsItemExistAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
	// 商品Id
	_itemId int64
	// 产品类型 101001005 代表普通推广，101001014代表销量明星
	_productId int64
}

// New
