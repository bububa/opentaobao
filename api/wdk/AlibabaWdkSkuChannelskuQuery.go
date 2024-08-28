package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuChannelskuQuery 查询渠道商品
// alibaba.wdk.sku.channelsku.query
//
// 查询渠道商品
func AlibabaWdkSkuChannelskuQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSkuChannelskuQueryAPIRequest, resp *wdk.AlibabaWdkSkuChannelskuQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
