package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvscreenAdminCommonOperation 一体机桌面通用接口
// yunos.tvscreen.admin.common.operation
//
// 一体机桌面通用接口
func YunosTvscreenAdminCommonOperation(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvscreenAdminCommonOperationAPIRequest, resp *tvupadmin.YunosTvscreenAdminCommonOperationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
