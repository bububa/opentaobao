package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest
迎客松查看小酷宝桌面坑位元数据列表 API请求
yunos.tvpubadmin.content.tableaudit.querychilddesktop

迎客松查看小酷宝桌面坑位元数据列表 */
type YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest struct {
	model.Params
	// 小酷宝桌面坑位查询参数
	_query string
}

// New
