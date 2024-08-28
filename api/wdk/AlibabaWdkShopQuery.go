package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkShopQuery 门店查询接口
// alibaba.wdk.shop.query
//
// 根据门店code查询门店信息
func AlibabaWdkShopQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkShopQueryAPIRequest, resp *wdk.AlibabaWdkShopQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
