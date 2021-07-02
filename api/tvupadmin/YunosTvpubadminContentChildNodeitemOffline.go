package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChildNodeitemOffline 少儿大厅类目内容下线接口
// yunos.tvpubadmin.content.child.nodeitem.offline
//
// 少儿大厅类目内容下线接口
func YunosTvpubadminContentChildNodeitemOffline(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildNodeitemOfflineAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentChildNodeitemOfflineAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentChildNodeitemOfflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
