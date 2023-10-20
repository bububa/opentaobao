package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentchildrecoitemoffline 下线少儿推荐内容接口
// yunos.tvpubadmin.content.child.recoitem.offline
//
// 下线少儿推荐内容接口
func Yunostvpubadmincontentchildrecoitemoffline(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentchildrecoitemofflineAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentchildrecoitemofflineAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentchildrecoitemofflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
