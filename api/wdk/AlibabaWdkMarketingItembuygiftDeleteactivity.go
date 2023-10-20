package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitembuygiftdeleteactivity 删除买赠活动
// alibaba.wdk.marketing.itembuygift.deleteactivity
//
// 删除买赠活动
func Alibabawdkmarketingitembuygiftdeleteactivity(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitembuygiftdeleteactivityAPIRequest, session string) (*wdk.AlibabawdkmarketingitembuygiftdeleteactivityAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitembuygiftdeleteactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
