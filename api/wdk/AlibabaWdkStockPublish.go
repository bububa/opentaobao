package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkStockPublish 五道口库存发布接口（针对外部渠道）
// alibaba.wdk.stock.publish
//
// 五道口库存发布接口（针对外部渠道）
func AlibabaWdkStockPublish(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkStockPublishAPIRequest, resp *wdk.AlibabaWdkStockPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
