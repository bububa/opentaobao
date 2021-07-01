package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenTradeUsersGetAPIRequest
获取奇门用户列表 API请求
taobao.qimen.trade.users.get

获取已开通奇门订单服务的用户列表 */
type TaobaoQimenTradeUsersGetAPIRequest struct {
	model.Params
	// 每页的数量
	_pageIndex int64
	// 页数
	_pageSize int64
}

// New
