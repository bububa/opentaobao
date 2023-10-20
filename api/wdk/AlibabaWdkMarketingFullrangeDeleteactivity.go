package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingfullrangedeleteactivity 全场活动删除活动接口
// alibaba.wdk.marketing.fullrange.deleteactivity
//
// 全场活动删除活动
func Alibabawdkmarketingfullrangedeleteactivity(clt *core.SDKClient, req *wdk.AlibabawdkmarketingfullrangedeleteactivityAPIRequest, session string) (*wdk.AlibabawdkmarketingfullrangedeleteactivityAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingfullrangedeleteactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
