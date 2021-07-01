package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest
运营位管理-联盟一体机下线运营位内容 API请求
yunos.tvpubadmin.content.tableaudit.offlinelauncheritem

运营位管理-联盟一体机下线运营位内容 */
type YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest struct {
	model.Params
	// 元数据主键id
	_id int64
	// 联盟：TV_OTT,一体机：TV_ALLINONE
	_terminalType string
}

// New
