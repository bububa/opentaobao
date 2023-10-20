package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolCreateactivity 添加商品池活动
// alibaba.hm.marketing.itempool.createactivity
//
// 添加商品池活动
func AlibabaHmMarketingItempoolCreateactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolCreateactivityAPIRequest, resp *wdk.AlibabaHmMarketingItempoolCreateactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
