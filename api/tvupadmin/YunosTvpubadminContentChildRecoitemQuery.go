package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChildRecoitemQuery 查询少儿大厅推荐内容列表
// yunos.tvpubadmin.content.child.recoitem.query
//
// 查询少儿大厅推荐内容列表
func YunosTvpubadminContentChildRecoitemQuery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildRecoitemQueryAPIRequest, resp *tvupadmin.YunosTvpubadminContentChildRecoitemQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
