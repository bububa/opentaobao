package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChildRecoitemOffline 下线少儿推荐内容接口
// yunos.tvpubadmin.content.child.recoitem.offline
//
// 下线少儿推荐内容接口
func YunosTvpubadminContentChildRecoitemOffline(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildRecoitemOfflineAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentChildRecoitemOfflineAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentChildRecoitemOfflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
