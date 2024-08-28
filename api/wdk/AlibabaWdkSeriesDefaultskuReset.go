package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSeriesDefaultskuReset 系列品商品变更-重置默认商品
// alibaba.wdk.series.defaultsku.reset
//
// 系列品商品变更-重置默认商品
func AlibabaWdkSeriesDefaultskuReset(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSeriesDefaultskuResetAPIRequest, resp *wdk.AlibabaWdkSeriesDefaultskuResetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
