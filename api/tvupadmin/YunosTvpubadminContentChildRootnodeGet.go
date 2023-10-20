package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentchildrootnodeget 获取少儿大厅根类目接口
// yunos.tvpubadmin.content.child.rootnode.get
//
// 通过此接口可获取少儿大厅根类目列表
func Yunostvpubadmincontentchildrootnodeget(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentchildrootnodegetAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentchildrootnodegetAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentchildrootnodegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
