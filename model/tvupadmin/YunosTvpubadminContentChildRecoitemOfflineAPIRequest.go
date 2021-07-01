package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentChildRecoitemOfflineAPIRequest
下线少儿推荐内容接口 API请求
yunos.tvpubadmin.content.child.recoitem.offline

下线少儿推荐内容接口 */
type YunosTvpubadminContentChildRecoitemOfflineAPIRequest struct {
	model.Params
	// 推荐内容ID
	_recItemId int64
}

// New
