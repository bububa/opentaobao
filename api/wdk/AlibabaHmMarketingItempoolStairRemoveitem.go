package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolStairRemoveitem 删除换购活动商品
// alibaba.hm.marketing.itempool.stair.removeitem
//
// 删除换购商品
func AlibabaHmMarketingItempoolStairRemoveitem(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolStairRemoveitemAPIRequest, resp *wdk.AlibabaHmMarketingItempoolStairRemoveitemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
