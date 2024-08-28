package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuFeature 商品标记接口
// alibaba.wdk.sku.feature
//
// 给淘鲜达商品属性之外的打标通用能力，满足商品一些特殊的需求，比如是否参加营销。
func AlibabaWdkSkuFeature(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSkuFeatureAPIRequest, resp *wdk.AlibabaWdkSkuFeatureAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
