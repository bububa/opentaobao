package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentappqueryapp 查询应用信息
// yunos.tvpubadmin.content.app.queryapp
//
// 查询应用信息
func Yunostvpubadmincontentappqueryapp(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentappqueryappAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentappqueryappAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentappqueryappAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
