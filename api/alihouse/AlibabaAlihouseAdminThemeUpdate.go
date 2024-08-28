package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseAdminThemeUpdate 云主题更新
// alibaba.alihouse.admin.theme.update
//
// 云主题更新
func AlibabaAlihouseAdminThemeUpdate(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseAdminThemeUpdateAPIRequest, resp *alihouse.AlibabaAlihouseAdminThemeUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
