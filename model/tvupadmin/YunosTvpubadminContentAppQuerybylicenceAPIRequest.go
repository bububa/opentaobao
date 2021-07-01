package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentAppQuerybylicenceAPIRequest
按牌照查询应用 API请求
yunos.tvpubadmin.content.app.querybylicence

按牌照查询应用 */
type YunosTvpubadminContentAppQuerybylicenceAPIRequest struct {
	model.Params
	// 查询条件
	_query string
}

// New
