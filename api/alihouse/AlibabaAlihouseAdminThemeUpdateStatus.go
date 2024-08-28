package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseAdminThemeUpdateStatus 云主题上下架+删除
// alibaba.alihouse.admin.theme.update.status
//
// 云主题上下架+删除
func AlibabaAlihouseAdminThemeUpdateStatus(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseAdminThemeUpdateStatusAPIRequest, resp *alihouse.AlibabaAlihouseAdminThemeUpdateStatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
