package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingBuygiftItemAddAsync 批量发布买赠商品
// alibaba.hm.marketing.buygift.item.add.async
//
// 批量发布买赠商品
func AlibabaHmMarketingBuygiftItemAddAsync(clt *core.SDKClient, req *wdk.AlibabaHmMarketingBuygiftItemAddAsyncAPIRequest, resp *wdk.AlibabaHmMarketingBuygiftItemAddAsyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
