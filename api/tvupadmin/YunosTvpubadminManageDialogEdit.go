package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageDialogEdit 编辑全局弹窗
// yunos.tvpubadmin.manage.dialog.edit
//
// 编辑全局弹窗
func YunosTvpubadminManageDialogEdit(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageDialogEditAPIRequest, resp *tvupadmin.YunosTvpubadminManageDialogEditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
