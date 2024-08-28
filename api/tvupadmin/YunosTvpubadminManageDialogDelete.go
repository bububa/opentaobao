package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageDialogDelete 删除全局弹窗
// yunos.tvpubadmin.manage.dialog.delete
//
// 删除全局弹窗
func YunosTvpubadminManageDialogDelete(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageDialogDeleteAPIRequest, resp *tvupadmin.YunosTvpubadminManageDialogDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
