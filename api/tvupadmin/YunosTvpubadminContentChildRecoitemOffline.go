package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChildRecoitemOffline 下线少儿推荐内容接口
// yunos.tvpubadmin.content.child.recoitem.offline
//
// 下线少儿推荐内容接口
func YunosTvpubadminContentChildRecoitemOffline(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildRecoitemOfflineAPIRequest, resp *tvupadmin.YunosTvpubadminContentChildRecoitemOfflineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
