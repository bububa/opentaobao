package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItempoolActivityDelete 删除商品池活动【同城零售】
// alibaba.retail.marketing.itempool.activity.delete
//
// 同城零售商品池活动删除
func AlibabaRetailMarketingItempoolActivityDelete(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItempoolActivityDeleteAPIRequest, resp *wdk.AlibabaRetailMarketingItempoolActivityDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
