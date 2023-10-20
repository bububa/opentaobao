package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolStairAdditem 商品池阶梯商品添加
// alibaba.hm.marketing.itempool.stair.additem
//
// 添加商品池阶梯商品
func AlibabaHmMarketingItempoolStairAdditem(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolStairAdditemAPIRequest, resp *wdk.AlibabaHmMarketingItempoolStairAdditemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
