package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingDiscountItemRemoveAsync 特价批量移除商品
// alibaba.hm.marketing.discount.item.remove.async
//
// 特价批量移除商品
func AlibabaHmMarketingDiscountItemRemoveAsync(clt *core.SDKClient, req *wdk.AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest, resp *wdk.AlibabaHmMarketingDiscountItemRemoveAsyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
