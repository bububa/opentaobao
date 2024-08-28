package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItempoolActivitySave 【同城零售】保存商品池活动
// alibaba.retail.marketing.itempool.activity.save
//
// 同城零售商品池活动保存
func AlibabaRetailMarketingItempoolActivitySave(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItempoolActivitySaveAPIRequest, resp *wdk.AlibabaRetailMarketingItempoolActivitySaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
