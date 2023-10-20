package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunitySubmittingText 聚划算社群动态文案下发接口
// alibaba.jhs.community.submitting.text
//
// 聚划算社群动态文案下发接口
func AlibabaJhsCommunitySubmittingText(clt *core.SDKClient, req *ju.AlibabaJhsCommunitySubmittingTextAPIRequest, resp *ju.AlibabaJhsCommunitySubmittingTextAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
