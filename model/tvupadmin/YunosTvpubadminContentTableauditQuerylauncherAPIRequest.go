package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentTableauditQuerylauncherAPIRequest
运营位管控-查询联盟一体机运营位元数据列表 API请求
yunos.tvpubadmin.content.tableaudit.querylauncher

运营位管控-查询联盟一体机运营位元数据列表 */
type YunosTvpubadminContentTableauditQuerylauncherAPIRequest struct {
	model.Params
	// 桌面坑位审核查询条件,json格式
	_tableAuditQuery string
}

// New
