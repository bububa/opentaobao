package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingFullrangeQueryitem 全场活动查询换购品
// alibaba.wdk.marketing.fullrange.queryitem
//
// 全场活动查询换购品
func AlibabaWdkMarketingFullrangeQueryitem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingFullrangeQueryitemAPIRequest, resp *wdk.AlibabaWdkMarketingFullrangeQueryitemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
