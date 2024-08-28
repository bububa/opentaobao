package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeMetroSync 地铁数据同步
// alibaba.alihouse.newhome.metro.sync
//
// 地铁数据同步
func AlibabaAlihouseNewhomeMetroSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeMetroSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeMetroSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
