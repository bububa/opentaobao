package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChildNodeitemOffline 少儿大厅类目内容下线接口
// yunos.tvpubadmin.content.child.nodeitem.offline
//
// 少儿大厅类目内容下线接口
func YunosTvpubadminContentChildNodeitemOffline(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildNodeitemOfflineAPIRequest, resp *tvupadmin.YunosTvpubadminContentChildNodeitemOfflineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
