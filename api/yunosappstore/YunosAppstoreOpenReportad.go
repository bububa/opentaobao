package yunosappstore

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosappstore"
)

// YunosAppstoreOpenReportad 外投广告上报接口
// yunos.appstore.open.reportad
//
// 外投广告回流上报接口
func YunosAppstoreOpenReportad(ctx context.Context, clt *core.SDKClient, req *yunosappstore.YunosAppstoreOpenReportadAPIRequest, resp *yunosappstore.YunosAppstoreOpenReportadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
