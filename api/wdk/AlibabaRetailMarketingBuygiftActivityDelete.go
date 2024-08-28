package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingBuygiftActivityDelete 删除单品买赠活动
// alibaba.retail.marketing.buygift.activity.delete
//
// 同城零售单品特价活动删除
func AlibabaRetailMarketingBuygiftActivityDelete(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest, resp *wdk.AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
