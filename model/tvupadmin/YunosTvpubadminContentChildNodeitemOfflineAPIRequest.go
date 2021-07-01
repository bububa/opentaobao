package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentChildNodeitemOfflineAPIRequest
少儿大厅类目内容下线接口 API请求
yunos.tvpubadmin.content.child.nodeitem.offline

少儿大厅类目内容下线接口 */
type YunosTvpubadminContentChildNodeitemOfflineAPIRequest struct {
	model.Params
	// 类目内容ID
	_contentId int64
}

// New
