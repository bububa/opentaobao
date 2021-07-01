package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallItemsQueryAPIRequest
批量获取商品列表 API请求
taobao.openmall.items.query

批量获取对联盟开放的商品列表。 */
type TaobaoOpenmallItemsQueryAPIRequest struct {
	model.Params
	// 已废弃，请勿使用
	_itemIds string
	// 第几页，默认：1
	_pageNo int64
	// 页大小，默认20，1~100
	_pageSize int64
	// 当不输入渠道商时，展示全网公有商品池；当输入渠道商的淘宝Nick时，展示该渠道私有供给商品列表
	_distributor string
}

// New
