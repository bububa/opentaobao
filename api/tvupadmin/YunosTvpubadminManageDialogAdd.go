package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageDialogAdd 新增全局弹窗
// yunos.tvpubadmin.manage.dialog.add
//
// 新增全局弹窗
func YunosTvpubadminManageDialogAdd(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageDialogAddAPIRequest, resp *tvupadmin.YunosTvpubadminManageDialogAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
