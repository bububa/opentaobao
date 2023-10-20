package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingopendarunfaactivityskusync 营销商品数据同步
// alibaba.wdk.marketing.open.darunfa.activity.sku.sync
//
// 大润发营销商品数据同步
func Alibabawdkmarketingopendarunfaactivityskusync(clt *core.SDKClient, req *wdk.AlibabawdkmarketingopendarunfaactivityskusyncAPIRequest, session string) (*wdk.AlibabawdkmarketingopendarunfaactivityskusyncAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingopendarunfaactivityskusyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
