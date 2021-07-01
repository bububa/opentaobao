package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeSpecialUsersQueryAPIRequest
专属下单标记信息查询 API请求
taobao.opentrade.special.users.query

专属下单标记信息查询 */
type TaobaoOpentradeSpecialUsersQueryAPIRequest struct {
	model.Params
	// 用户openId列表，多个以逗号(,)分割
	_openUserIds []string
	// 分页参数，每页大小
	_pageSize int64
	// 商品ID
	_itemId int64
	// 商品SKU ID，不存在传0
	_skuId int64
	// 用户状态
	_status string
	// 分页参数，当前页，以0开始
	_pageIndex int64
}

// New
