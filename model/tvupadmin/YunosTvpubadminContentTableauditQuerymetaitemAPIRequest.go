package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentTableauditQuerymetaitemAPIRequest
运营位管控-查询魔盒运营位元数据列表 API请求
yunos.tvpubadmin.content.tableaudit.querymetaitem

运营位管控-查询魔盒运营位元数据列表 */
type YunosTvpubadminContentTableauditQuerymetaitemAPIRequest struct {
	model.Params
	// 查询条件，json格式
	_tableAuditQueryBo string
}

// New
