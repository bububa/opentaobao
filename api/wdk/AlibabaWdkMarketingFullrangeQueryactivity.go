package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingFullrangeQueryactivity 全场活动查询活动
// alibaba.wdk.marketing.fullrange.queryactivity
//
// 全场活动查询活动
func AlibabaWdkMarketingFullrangeQueryactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingFullrangeQueryactivityAPIRequest, session string) (*wdk.AlibabaWdkMarketingFullrangeQueryactivityAPIResponse, error) {
	var resp wdk.AlibabaWdkMarketingFullrangeQueryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
