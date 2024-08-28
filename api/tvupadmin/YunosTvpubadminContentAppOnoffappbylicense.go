package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentAppOnoffappbylicense 应用上下架操作
// yunos.tvpubadmin.content.app.onoffappbylicense
//
// 应用上下架操作
func YunosTvpubadminContentAppOnoffappbylicense(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentAppOnoffappbylicenseAPIRequest, resp *tvupadmin.YunosTvpubadminContentAppOnoffappbylicenseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
