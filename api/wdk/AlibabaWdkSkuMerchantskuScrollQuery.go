package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuMerchantskuScrollQuery 商家商品批量查询接口
// alibaba.wdk.sku.merchantsku.scroll.query
//
// 提供主档商品数据接口查询
func AlibabaWdkSkuMerchantskuScrollQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSkuMerchantskuScrollQueryAPIRequest, resp *wdk.AlibabaWdkSkuMerchantskuScrollQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
