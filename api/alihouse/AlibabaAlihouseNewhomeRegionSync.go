package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeRegionSync 城区数据同步
// alibaba.alihouse.newhome.region.sync
//
// 城区数据同步
func AlibabaAlihouseNewhomeRegionSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeRegionSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeRegionSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
