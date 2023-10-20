package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChildRootnodeGet 获取少儿大厅根类目接口
// yunos.tvpubadmin.content.child.rootnode.get
//
// 通过此接口可获取少儿大厅根类目列表
func YunosTvpubadminContentChildRootnodeGet(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildRootnodeGetAPIRequest, resp *tvupadmin.YunosTvpubadminContentChildRootnodeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
