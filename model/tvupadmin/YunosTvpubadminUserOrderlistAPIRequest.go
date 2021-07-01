package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminUserOrderlistAPIRequest
获取用户订单列表 API请求
yunos.tvpubadmin.user.orderlist

获取用户订单列表 */
type YunosTvpubadminUserOrderlistAPIRequest struct {
	model.Params
	// 用户ID
	_uid int64
	// 开始时间
	_createTimeStartStr string
	// 结束时间
	_createTimeEndStr string
	// 牌照方
	_license int64
	// 页码值
	_pageNo int64
	// 每页行数
	_pageSize int64
}

// New
