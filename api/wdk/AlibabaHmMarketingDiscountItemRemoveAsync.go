package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingDiscountItemRemoveAsync 特价批量移除商品
// alibaba.hm.marketing.discount.item.remove.async
//
// 特价批量移除商品
func AlibabaHmMarketingDiscountItemRemoveAsync(clt *core.SDKClient, req *wdk.AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest, session string) (*wdk.AlibabaHmMarketingDiscountItemRemoveAsyncAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingDiscountItemRemoveAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
