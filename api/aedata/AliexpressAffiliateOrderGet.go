package aedata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedata"
)

// AliexpressAffiliateOrderGet AE流量订单详情获取API
// aliexpress.affiliate.order.get
//
// 联盟推广订单效果获取API
func AliexpressAffiliateOrderGet(clt *core.SDKClient, req *aedata.AliexpressAffiliateOrderGetAPIRequest, resp *aedata.AliexpressAffiliateOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
