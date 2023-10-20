package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunityActivityDetails 社群活动详情
// alibaba.jhs.community.activity.details
//
// 社群活动详情
func AlibabaJhsCommunityActivityDetails(clt *core.SDKClient, req *ju.AlibabaJhsCommunityActivityDetailsAPIRequest, resp *ju.AlibabaJhsCommunityActivityDetailsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
