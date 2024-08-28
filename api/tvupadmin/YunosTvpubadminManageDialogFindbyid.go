package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageDialogFindbyid 根据id查询全局弹窗
// yunos.tvpubadmin.manage.dialog.findbyid
//
// 根据id查询全局弹窗
func YunosTvpubadminManageDialogFindbyid(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageDialogFindbyidAPIRequest, resp *tvupadmin.YunosTvpubadminManageDialogFindbyidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
