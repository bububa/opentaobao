package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingBuygiftActivityUpdate 更新单品买赠活动
// alibaba.retail.marketing.buygift.activity.update
//
// 同城零售单品买赠活动更新
func AlibabaRetailMarketingBuygiftActivityUpdate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest, resp *wdk.AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
