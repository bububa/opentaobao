package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminUserSuggestAPIRequest
获取关联账户列表 API请求
yunos.tvpubadmin.user.suggest

获取关联账户列表 */
type YunosTvpubadminUserSuggestAPIRequest struct {
	model.Params
	// 关键词
	_keyword string
}

// New
