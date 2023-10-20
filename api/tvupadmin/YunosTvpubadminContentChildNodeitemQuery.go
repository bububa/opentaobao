package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChildNodeitemQuery 查询少儿大厅类目内容
// yunos.tvpubadmin.content.child.nodeitem.query
//
// 查询少儿大厅类目内容信息
func YunosTvpubadminContentChildNodeitemQuery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildNodeitemQueryAPIRequest, resp *tvupadmin.YunosTvpubadminContentChildNodeitemQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
