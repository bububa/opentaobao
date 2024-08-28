package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkBmPaiyangSupplyPublish 派样商品库存变更同步接口
// alibaba.wdk.bm.paiyang.supply.publish
//
// 淘鲜达接入第三方进行派样，第三方同步大仓和门店的库存变更信息。
func AlibabaWdkBmPaiyangSupplyPublish(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkBmPaiyangSupplyPublishAPIRequest, resp *wdk.AlibabaWdkBmPaiyangSupplyPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
