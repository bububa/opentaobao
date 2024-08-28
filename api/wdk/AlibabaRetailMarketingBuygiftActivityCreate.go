package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingBuygiftActivityCreate 创建买赠活动
// alibaba.retail.marketing.buygift.activity.create
//
// 同城供给买赠活动创建
func AlibabaRetailMarketingBuygiftActivityCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftActivityCreateAPIRequest, resp *wdk.AlibabaRetailMarketingBuygiftActivityCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
