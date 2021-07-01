package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentAppQueryappAPIRequest
查询应用信息 API请求
yunos.tvpubadmin.content.app.queryapp

查询应用信息 */
type YunosTvpubadminContentAppQueryappAPIRequest struct {
	model.Params
	// 查询条件
	_query string
}

// New
