package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolStairQueryitem 换购商品查询
// alibaba.hm.marketing.itempool.stair.queryitem
//
// 换购商品查询
func AlibabaHmMarketingItempoolStairQueryitem(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolStairQueryitemAPIRequest, resp *wdk.AlibabaHmMarketingItempoolStairQueryitemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
