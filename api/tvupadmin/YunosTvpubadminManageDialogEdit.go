package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageDialogEdit 编辑全局弹窗
// yunos.tvpubadmin.manage.dialog.edit
//
// 编辑全局弹窗
func YunosTvpubadminManageDialogEdit(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageDialogEditAPIRequest, resp *tvupadmin.YunosTvpubadminManageDialogEditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
