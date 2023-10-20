package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentShowEdit 媒资节目信息修改
// yunos.tvpubadmin.content.show.edit
//
// 供迎客松修改媒资节目信息
func YunosTvpubadminContentShowEdit(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentShowEditAPIRequest, resp *tvupadmin.YunosTvpubadminContentShowEditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
