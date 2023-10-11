package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingFullrangeQueryactivity 全场活动查询活动
// alibaba.hm.marketing.fullrange.queryactivity
//
// 全场活动查询活动
func AlibabaHmMarketingFullrangeQueryactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingFullrangeQueryactivityAPIRequest, session string) (*wdk.AlibabaHmMarketingFullrangeQueryactivityAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingFullrangeQueryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
