package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItempoolActivityQuery 查询商品池活动【同城零售】
// alibaba.retail.marketing.itempool.activity.query
//
// 查询商品池活动【同城零售】
func AlibabaRetailMarketingItempoolActivityQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItempoolActivityQueryAPIRequest, resp *wdk.AlibabaRetailMarketingItempoolActivityQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
