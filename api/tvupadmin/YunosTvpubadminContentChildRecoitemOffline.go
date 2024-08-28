package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChildRecoitemOffline 下线少儿推荐内容接口
// yunos.tvpubadmin.content.child.recoitem.offline
//
// 下线少儿推荐内容接口
func YunosTvpubadminContentChildRecoitemOffline(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildRecoitemOfflineAPIRequest, resp *tvupadmin.YunosTvpubadminContentChildRecoitemOfflineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
