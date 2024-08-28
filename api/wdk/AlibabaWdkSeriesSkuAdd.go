package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSeriesSkuAdd 系列品商品变更-添加商品
// alibaba.wdk.series.sku.add
//
// 系列品商品变更-添加商品
func AlibabaWdkSeriesSkuAdd(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSeriesSkuAddAPIRequest, resp *wdk.AlibabaWdkSeriesSkuAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
