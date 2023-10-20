package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItembuygiftCreateactivity 创建买赠活动
// alibaba.hm.marketing.itembuygift.createactivity
//
// 创建买赠活动
func AlibabaHmMarketingItembuygiftCreateactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItembuygiftCreateactivityAPIRequest, session string) (*wdk.AlibabaHmMarketingItembuygiftCreateactivityAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItembuygiftCreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
