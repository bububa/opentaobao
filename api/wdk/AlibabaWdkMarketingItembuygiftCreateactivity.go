package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitembuygiftcreateactivity 创建买赠活动
// alibaba.wdk.marketing.itembuygift.createactivity
//
// 创建买赠活动
func Alibabawdkmarketingitembuygiftcreateactivity(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitembuygiftcreateactivityAPIRequest, session string) (*wdk.AlibabawdkmarketingitembuygiftcreateactivityAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitembuygiftcreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
