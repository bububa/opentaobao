package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingfullrangecreateactivity 创建全场活动
// alibaba.wdk.marketing.fullrange.createactivity
//
// 创建全场活动
func Alibabawdkmarketingfullrangecreateactivity(clt *core.SDKClient, req *wdk.AlibabawdkmarketingfullrangecreateactivityAPIRequest, session string) (*wdk.AlibabawdkmarketingfullrangecreateactivityAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingfullrangecreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
