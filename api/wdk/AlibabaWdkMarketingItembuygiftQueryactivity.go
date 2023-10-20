package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitembuygiftqueryactivity 查询买赠活动
// alibaba.wdk.marketing.itembuygift.queryactivity
//
// 查询买赠活动
func Alibabawdkmarketingitembuygiftqueryactivity(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitembuygiftqueryactivityAPIRequest, session string) (*wdk.AlibabawdkmarketingitembuygiftqueryactivityAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitembuygiftqueryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
