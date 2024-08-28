package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuChannelskuAdd 新增渠道商品
// alibaba.wdk.sku.channelsku.add
//
// 盒马帮1期新增渠道商品
func AlibabaWdkSkuChannelskuAdd(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSkuChannelskuAddAPIRequest, resp *wdk.AlibabaWdkSkuChannelskuAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
