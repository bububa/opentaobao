package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingfullrangequeryactivity 全场活动查询活动
// alibaba.wdk.marketing.fullrange.queryactivity
//
// 全场活动查询活动
func Alibabawdkmarketingfullrangequeryactivity(clt *core.SDKClient, req *wdk.AlibabawdkmarketingfullrangequeryactivityAPIRequest, session string) (*wdk.AlibabawdkmarketingfullrangequeryactivityAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingfullrangequeryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
