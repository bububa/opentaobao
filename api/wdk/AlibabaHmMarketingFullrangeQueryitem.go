package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingFullrangeQueryitem 全场活动查询换购品
// alibaba.hm.marketing.fullrange.queryitem
//
// 全场活动查询换购品
func AlibabaHmMarketingFullrangeQueryitem(clt *core.SDKClient, req *wdk.AlibabaHmMarketingFullrangeQueryitemAPIRequest, session string) (*wdk.AlibabaHmMarketingFullrangeQueryitemAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingFullrangeQueryitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
