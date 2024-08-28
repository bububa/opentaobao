package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolExcludeskucode 商品池排除商品【品类优惠使用】
// alibaba.hm.marketing.itempool.excludeskucode
//
// 品类优惠新增排除池
func AlibabaHmMarketingItempoolExcludeskucode(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolExcludeskucodeAPIRequest, resp *wdk.AlibabaHmMarketingItempoolExcludeskucodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
