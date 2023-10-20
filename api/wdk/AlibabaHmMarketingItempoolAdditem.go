package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolAdditem 增加商品池里面的商品
// alibaba.hm.marketing.itempool.additem
//
// 增加商品池里面的商品
func AlibabaHmMarketingItempoolAdditem(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolAdditemAPIRequest, resp *wdk.AlibabaHmMarketingItempoolAdditemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
