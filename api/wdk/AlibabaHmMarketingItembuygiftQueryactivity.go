package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItembuygiftQueryactivity 查询买赠活动
// alibaba.hm.marketing.itembuygift.queryactivity
//
// 查询买赠活动
func AlibabaHmMarketingItembuygiftQueryactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItembuygiftQueryactivityAPIRequest, session string) (*wdk.AlibabaHmMarketingItembuygiftQueryactivityAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItembuygiftQueryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
