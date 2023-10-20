package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentAppQueryapp 查询应用信息
// yunos.tvpubadmin.content.app.queryapp
//
// 查询应用信息
func YunosTvpubadminContentAppQueryapp(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentAppQueryappAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentAppQueryappAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentAppQueryappAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
