package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItembuygiftDeleteactivity 删除买赠活动
// alibaba.hm.marketing.itembuygift.deleteactivity
//
// 删除买赠活动
func AlibabaHmMarketingItembuygiftDeleteactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest, session string) (*wdk.AlibabaHmMarketingItembuygiftDeleteactivityAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItembuygiftDeleteactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
