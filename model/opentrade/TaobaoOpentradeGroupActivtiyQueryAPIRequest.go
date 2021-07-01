package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeGroupActivtiyQueryAPIRequest
组团活动信息查询 API请求
taobao.opentrade.group.activtiy.query

组团购场景下，团购活动信息查询 */
type TaobaoOpentradeGroupActivtiyQueryAPIRequest struct {
	model.Params
	// 分页参数，每页大小
	_pageSize int64
	// 商品ID
	_itemId int64
	// 分页参数，当前页，以0开始
	_pageIndex int64
	// 组团活动id
	_groupActivityId int64
}

// New
