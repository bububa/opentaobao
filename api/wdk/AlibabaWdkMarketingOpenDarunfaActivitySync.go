package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingopendarunfaactivitysync 活动数据同步
// alibaba.wdk.marketing.open.darunfa.activity.sync
//
// 大润发活动数据同步
func Alibabawdkmarketingopendarunfaactivitysync(clt *core.SDKClient, req *wdk.AlibabawdkmarketingopendarunfaactivitysyncAPIRequest, session string) (*wdk.AlibabawdkmarketingopendarunfaactivitysyncAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingopendarunfaactivitysyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
