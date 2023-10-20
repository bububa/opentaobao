package lstmarketing

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstmarketing"
)

// AlibabaLstMarketingQuerybyorderid 根据订单查询营销信息
// alibaba.lst.marketing.querybyorderid
//
// 根据订单查询营销信息
func AlibabaLstMarketingQuerybyorderid(clt *core.SDKClient, req *lstmarketing.AlibabaLstMarketingQuerybyorderidAPIRequest, resp *lstmarketing.AlibabaLstMarketingQuerybyorderidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
