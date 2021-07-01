package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJushitaJdpUsersGetAPIRequest
获取开通的订单同步服务的用户 API请求
taobao.jushita.jdp.users.get

获取开通的订单同步服务的用户，含有rds的路由关系 */
type TaobaoJushitaJdpUsersGetAPIRequest struct {
	model.Params
	// 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
	_startModified string
	// 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
	_endModified string
	// 每页记录数，默认200
	_pageSize int64
	// 当前页数
	_pageNo int64
	// 如果传了user_id表示单条查询
	_userId int64
}

// New
