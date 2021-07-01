package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeGroupQueryAPIRequest
组团购场景查询商品组团详情 API请求
taobao.opentrade.group.query

组团购场景下，查询商品开团详情 */
type TaobaoOpentradeGroupQueryAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 用户openId
	_openUserId string
	// 0 返回未成团列表，1 返回已成团列表
	_orderBy int64
	// 页数
	_pageIndex int64
	// 每页展示条数，不能超过100
	_pageSize int64
	// 组团活动id
	_groupActivityId int64
	// 是否返回已过期的团，true 返回，false 不返回
	_withExpire bool
}

// New
