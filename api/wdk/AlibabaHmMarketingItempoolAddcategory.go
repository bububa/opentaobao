package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolAddcategory 增加商品池里面的类目
// alibaba.hm.marketing.itempool.addcategory
//
// 增加商品池里面的类目
func AlibabaHmMarketingItempoolAddcategory(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolAddcategoryAPIRequest, resp *wdk.AlibabaHmMarketingItempoolAddcategoryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
