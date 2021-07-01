package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminUserRightsAPIRequest
获取用户权益 API请求
yunos.tvpubadmin.user.rights

获取用户权益 */
type YunosTvpubadminUserRightsAPIRequest struct {
	model.Params
	// 用户ID
	_uid int64
	// 页码值
	_pageNo int64
	// 每页行数
	_pageSize int64
}

// New
