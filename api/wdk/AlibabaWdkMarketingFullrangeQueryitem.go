package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingfullrangequeryitem 全场活动查询换购品
// alibaba.wdk.marketing.fullrange.queryitem
//
// 全场活动查询换购品
func Alibabawdkmarketingfullrangequeryitem(clt *core.SDKClient, req *wdk.AlibabawdkmarketingfullrangequeryitemAPIRequest, session string) (*wdk.AlibabawdkmarketingfullrangequeryitemAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingfullrangequeryitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
