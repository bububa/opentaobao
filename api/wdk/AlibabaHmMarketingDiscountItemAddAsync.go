package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingDiscountItemAddAsync 特价批量新增商品
// alibaba.hm.marketing.discount.item.add.async
//
// 新分组模型下新增商品
func AlibabaHmMarketingDiscountItemAddAsync(clt *core.SDKClient, req *wdk.AlibabaHmMarketingDiscountItemAddAsyncAPIRequest, session string) (*wdk.AlibabaHmMarketingDiscountItemAddAsyncAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingDiscountItemAddAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
